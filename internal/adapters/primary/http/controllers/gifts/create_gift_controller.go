package gift_controllers

import (
	presenters "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/presenters/gifts"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
	create_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/create"
	"github.com/gin-gonic/gin"
)

type CreateGiftController struct {
	CreateGiftUC create_gift.CreateGiftUC
	Presenter    presenters.CreateGiftPresenter
}

func BuildSignUpController(uc *create_gift.CreateGiftUC, ptr *presenters.CreateGiftPresenter) *CreateGiftController {
	return &CreateGiftController{CreateGiftUC: *uc, Presenter: *ptr}
}

func (ctrl *CreateGiftController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)

	amount, ok := inputMap["amount"].(float64)
	if !ok {
		amount = 0
	}

	input := create_gift.CreateGiftInputDTO{
		Type:       domain.GiftType(inputMap["type"].(string)),
		CPF:        inputMap["cpf"].(string),
		Amount:     amount,
		PixKeyType: inputMap["pix_key_type"].(string),
		PixKey:     inputMap["pix_key"].(string),
		Message:    inputMap["message"].(string),
	}

	result, err := ctrl.CreateGiftUC.Execute(input)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
