package common_ptr

import (
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type AuthorizedGiftFormatted struct {
	ID             string `json:"id" example:"123" description:"O ID único da autorização"`
	GiftID         string `json:"gift_id" example:"123" description:"O ID único do presente"`
	ExpirationDate string `json:"expiration_date" example:"2022-01-01T00:00:00Z" description:"A data e hora de expiração da autorização"`
	CreatedAt      string `json:"created_at" example:"2022-01-01T00:00:00Z" description:"A data e hora de criação do presente"`
	UpdatedAt      string `json:"updated_at" example:"2022-01-01T01:00:00Z" description:"A data e hora da última atualização do presente"`
}

type AuthorizedGiftPresenter struct{}

func BuildAuthorizedGiftPresenter() *AuthorizedGiftPresenter {
	return &AuthorizedGiftPresenter{}
}

func (ptr *AuthorizedGiftPresenter) Format(ag domain.AuthorizedGift) AuthorizedGiftFormatted {
	agf := AuthorizedGiftFormatted{
		ID:        ag.ID,
		GiftID:    ag.GiftID,
		CreatedAt: ag.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: ag.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	if ag.ExpirationDate != nil {
		agf.ExpirationDate = ag.ExpirationDate.Format("2006-01-02T15:04:05Z")
	} else {
		agf.ExpirationDate = ""
	}

	return agf
}

func (ptr *AuthorizedGiftPresenter) FormatList(ags []domain.AuthorizedGift) []AuthorizedGiftFormatted {
	var agsf []AuthorizedGiftFormatted = make([]AuthorizedGiftFormatted, 0, len(ags))

	for _, ag := range ags {
		agsf = append(agsf, ptr.Format(ag))
	}

	return agsf
}
