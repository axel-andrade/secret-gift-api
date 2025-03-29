package common_ptr

import (
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type GiftFormatted struct {
	ID         string            `json:"id" example:"123" description:"O ID único do presente"`
	Type       domain.GiftType   `json:"type" example:"pix" description:"O tipo do presente"`
	CPF        string            `json:"cpf" example:"12345678900" description:"O CPF do destinatário"`
	Amount     float64           `json:"amount" example:"100.50" description:"O valor do presente"`
	PixKeyType string            `json:"pix_key_type" example:"email" description:"O tipo da chave Pix"`
	PixKey     string            `json:"pix_key" example:"user@example.com" description:"A chave Pix do destinatário"`
	Message    string            `json:"message" example:"Feliz aniversário!" description:"Mensagem do presente"`
	Status     domain.GiftStatus `json:"status" example:"pending" description:"O status do presente"`
	CreatedAt  string            `json:"created_at" example:"2022-01-01T00:00:00Z" description:"A data e hora de criação do presente"`
	UpdatedAt  string            `json:"updated_at" example:"2022-01-01T01:00:00Z" description:"A data e hora da última atualização do presente"`
}

type GiftPresenter struct{}

func BuildGiftPresenter() *GiftPresenter {
	return &GiftPresenter{}
}

func (ptr *GiftPresenter) Format(gift domain.Gift) GiftFormatted {
	return GiftFormatted{
		ID:         gift.ID,
		Type:       gift.Type,
		CPF:        gift.CPF,
		Amount:     gift.Amount,
		PixKeyType: gift.PixKeyType,
		PixKey:     gift.PixKey,
		Message:    gift.Message,
		Status:     gift.Status,
		CreatedAt:  gift.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:  gift.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func (ptr *GiftPresenter) FormatList(gifts []domain.Gift) []GiftFormatted {
	var giftsFormatted []GiftFormatted = make([]GiftFormatted, 0, len(gifts))

	for _, gift := range gifts {
		giftsFormatted = append(giftsFormatted, ptr.Format(gift))
	}

	return giftsFormatted
}
