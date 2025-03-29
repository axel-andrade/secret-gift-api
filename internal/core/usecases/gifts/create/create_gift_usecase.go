package create_gift

import (
	"log"

	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type CreateGiftUC struct {
	Gateway CreateGiftGateway
}

func BuildCreateGiftUC(g CreateGiftGateway) *CreateGiftUC {
	return &CreateGiftUC{g}
}

func (bs *CreateGiftUC) Execute(input CreateGiftInputDTO) (*CreateGiftOutputDTO, error) {
	log.Println("info: building gift entity")
	gift, _ := domain.NewGift(
		input.Type,
		input.Amount,
		input.CPF,
		input.PixKeyType,
		input.PixKey,
		input.Message,
	)

	createdGift, err := bs.Gateway.CreateGift(gift)
	if err != nil {
		return nil, err
	}

	log.Println("info: gift created with success")

	return &CreateGiftOutputDTO{Gift: createdGift}, nil
}
