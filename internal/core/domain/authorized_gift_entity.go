package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AuthorizedGift struct {
	Base
	GiftID         uuid.UUID  `json:"gift_id"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
}

func NewAuthorizedGift(giftID uuid.UUID, expirationDate *time.Time) (*AuthorizedGift, error) {
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
	if ag.GiftID == uuid.Nil {
		return errors.New("gift_id is required")
	}

	return nil
}
