package authorized_gift_controllers

import (
	"time"

	presenters "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/presenters/gifts"
	authorize_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/authorize"
	"github.com/gin-gonic/gin"
)

type AuthorizeGiftController struct {
	AuthorizedGiftUC authorize_gift.AuthorizeGiftUC
	Presenter        presenters.AuthorizeGiftPresenter
}

func BuildAuthorizeGiftController(uc *authorize_gift.AuthorizeGiftUC, ptr *presenters.AuthorizeGiftPresenter) *AuthorizeGiftController {
	return &AuthorizeGiftController{AuthorizedGiftUC: *uc, Presenter: *ptr}
}

func (ctrl *AuthorizeGiftController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)

	input := authorize_gift.AuthorizeGiftInput{
		GiftID: inputMap["gift_id"].(string),
		ExpirationDate: func() *time.Time {
			if dateStr, ok := inputMap["expiration_date"].(string); ok {
				if parsedDate, err := time.Parse(time.RFC3339, dateStr); err == nil {
					return &parsedDate
				}
			}
			return nil
		}(),
	}

	result, err := ctrl.AuthorizedGiftUC.Execute(input)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
