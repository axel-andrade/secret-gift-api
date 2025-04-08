package authorize_gift

import (
	"time"

	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type AuthorizeGiftGateway interface {
	GetGift(id string) (*domain.Gift, error)
	GetAuthorizedGift(giftId string) (*domain.AuthorizedGift, error)
	CreateAuthorizedGift(ag *domain.AuthorizedGift) (*domain.AuthorizedGift, error)
}

type AuthorizeGiftInput struct {
	GiftID         string     `json:"gift_id"`
	ExpirationDate *time.Time `json:"expiration_date"`
}

type AuthorizeGiftOutput struct {
	AuthorizedGift *domain.AuthorizedGift
}
