package authorize_gift

import (
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type AuthorizeGiftGateway interface {
	CreateGift(gift *domain.Gift) (*domain.Gift, error)
}

type CreateGiftInputDTO struct {
	Type       domain.GiftType `json:"type"`
	CPF        string          `json:"cpf"`
	Amount     float64         `json:"amount"`
	PixKeyType string          `json:"pix_key_type"`
	PixKey     string          `json:"pix_key"`
	Message    string          `json:"message"`
}

type CreateGiftOutputDTO struct {
	Gift *domain.Gift
}
