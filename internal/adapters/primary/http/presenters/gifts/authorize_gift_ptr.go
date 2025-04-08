package gift_presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/presenters/common"
	err_msg "github.com/axel-andrade/secret-gift-api/internal/core/domain/constants/errors"
	authorize_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/authorize"
)

type AuthorizeGiftPresenter struct {
	authorizedGiftPtr common_ptr.AuthorizedGiftPresenter
}

func BuildAuthorizeGiftPresenter() *AuthorizeGiftPresenter {
	return &AuthorizeGiftPresenter{}
}

func (p *AuthorizeGiftPresenter) Show(result *authorize_gift.AuthorizeGiftOutput, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	data := p.authorizedGiftPtr.Format(*result.AuthorizedGift)

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: data}
}

func (p *AuthorizeGiftPresenter) formatError(err error) common_adapters.OutputPort {
	if err.Error() == err_msg.GIFT_NOT_FOUND {
		return common_adapters.OutputPort{StatusCode: http.StatusNotFound, Data: common_adapters.ErrorMessage{Message: err_msg.GIFT_NOT_FOUND}}
	}

	if err.Error() == err_msg.AUTHORIZED_GIFT_ALREADY_EXISTS {
		return common_adapters.OutputPort{StatusCode: http.StatusConflict, Data: common_adapters.ErrorMessage{Message: err_msg.AUTHORIZED_GIFT_ALREADY_EXISTS}}
	}

	return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
}
