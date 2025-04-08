package authorize_gift

import (
	"fmt"
	"log"

	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
	err_msg "github.com/axel-andrade/secret-gift-api/internal/core/domain/constants/errors"
)

type AuthorizeGiftUC struct {
	Gateway AuthorizeGiftGateway
}

func BuildAuthorizeGiftUC(g AuthorizeGiftGateway) *AuthorizeGiftUC {
	return &AuthorizeGiftUC{g}
}

func (bs *AuthorizeGiftUC) Execute(input AuthorizeGiftInput) (*AuthorizeGiftOutput, error) {
	log.Println("info: getting gift")
	gift, err := bs.Gateway.GetGift(input.GiftID)

	if err != nil {
		log.Println("error: error getting gift")
		return nil, err
	}

	if gift == nil {
		log.Println("error: gift not found")
		return nil, fmt.Errorf(err_msg.GIFT_NOT_FOUND)
	}

	log.Println(("info: getting authorized gift"))
	currAg, _ := bs.Gateway.GetAuthorizedGift(input.GiftID)

	if currAg != nil {
		log.Println("error: authorized gift already exists")
		return nil, fmt.Errorf(err_msg.AUTHORIZED_GIFT_ALREADY_EXISTS)
	}

	log.Println("info: building authorized gift entity")
	newAg, _ := domain.NewAuthorizedGift(
		input.GiftID,
		input.ExpirationDate,
	)

	log.Println("info: creating authorized gift entity")
	createdAg, err := bs.Gateway.CreateAuthorizedGift(newAg)
	if err != nil {
		return nil, err
	}

	log.Println("info: authorized gift created with success")

	return &AuthorizeGiftOutput{AuthorizedGift: createdAg}, nil
}
