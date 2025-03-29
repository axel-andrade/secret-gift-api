package domain

import (
	"fmt"

	vo "github.com/axel-andrade/secret-gift-api/internal/core/domain/value_objects"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Base
	Username         vo.Email    `json:"username"`
	Password         vo.Password `json:"-"`
	Role             Role        `json:"role"`
	RefreshToken     *string     `json:"refresh_token"`
	IsEmailConfirmed bool        `json:"is_email_confirmed"`
}

func BuildUser(
	email, password string,
	role Role,
	options ...func(*User),
) (*User, error) {
	u := &User{
		Username:         vo.Email{Value: email},
		Password:         vo.Password{Value: password},
		Role:             role,
		IsEmailConfirmed: false,
	}

	for _, option := range options {
		option(u)
	}

	if err := u.validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) validate() error {
	if err := u.Username.Validate(); err != nil {
		return err
	}

	if err := u.Password.Validate(); err != nil {
		return err
	}

	if u.Role != RoleAdmin && u.Role != RoleUser {
		return fmt.Errorf("invalid role")
	}

	return nil
}
