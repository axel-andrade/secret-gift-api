package domain

import (
	"errors"
	"time"
)

type AuthorizedGift struct {
	Base
	GiftID         string     `json:"gift_id"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
}

func NewAuthorizedGift(giftID string, expirationDate *time.Time) (*AuthorizedGift, error) {
	ag := &AuthorizedGift{
		GiftID:         giftID,
		ExpirationDate: expirationDate,
	}

	if err := ag.validate(); err != nil {
		return nil, err
	}

	return ag, nil
}

func (ag *AuthorizedGift) validate() error {
	if ag.GiftID == "" {
		return errors.New("gift_id is required")
	}

	return nil
}
