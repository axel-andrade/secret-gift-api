package routes

import (
	"github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/secret-gift-api/internal/infra/bootstrap"
	"github.com/gin-gonic/gin"
)

func configureAuthorizedGiftsRoutes(r *gin.RouterGroup, d *bootstrap.Dependencies) {
	ag := r.Group("authorized-gifts")
	{
		ag.POST("", middlewares.ValidateRequest("authorized_gifts/authorize_gift"), d.AuthorizeGiftController.Handle)
	}
}
