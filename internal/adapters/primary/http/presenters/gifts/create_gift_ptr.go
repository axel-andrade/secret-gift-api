package gift_presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/presenters/common"
	err_msg "github.com/axel-andrade/secret-gift-api/internal/core/domain/constants/errors"
	create_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/create"
)

type CreateGiftPresenter struct {
	giftPtr common_ptr.GiftPresenter
}

func BuildCreateGiftPresenter() *CreateGiftPresenter {
	return &CreateGiftPresenter{}
}

func (p *CreateGiftPresenter) Show(result *create_gift.CreateGiftOutputDTO, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError()
	}

	data := p.giftPtr.Format(*result.Gift)

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: data}
}

func (p *CreateGiftPresenter) formatError() common_adapters.OutputPort {
	return common_adapters.OutputPort{StatusCode: http.StatusInternalServerError, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
}
