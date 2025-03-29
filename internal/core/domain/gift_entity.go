package domain

import (
	"errors"
)

type GiftStatus string

type GiftType string

const (
	GiftStatusPending   GiftStatus = "pending"
	GiftStatusSent      GiftStatus = "sent"
	GiftStatusConfirmed GiftStatus = "authorized"
	GiftStatusRejected  GiftStatus = "rejected"
)

const (
	GiftTypePix     GiftType = "pix"
	GiftTypeMessage GiftType = "message"
)

type Gift struct {
	Base
	Type       GiftType   `json:"type"`
	CPF        string     `json:"cpf"`
	Amount     float64    `json:"amount"`
	PixKeyType string     `json:"pix_key_type"`
	PixKey     string     `json:"pix_key"`
	Message    string     `json:"message"`
	Status     GiftStatus `json:"status"`
}

func NewGift(giftType GiftType, amount float64, cpf string, pixKeyType, pixKey, message string) (*Gift, error) {
	g := &Gift{
		Type:       giftType,
		CPF:        cpf,
		Amount:     amount,
		PixKeyType: pixKeyType,
		PixKey:     pixKey,
		Message:    message,
		Status:     GiftStatusPending,
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Gift) validate() error {
	if g.CPF == "" {
		return errors.New("cpf is required")
	}

	if g.Status == "" {
		return errors.New("status is required")
	}

	return nil
}
