package routes

import (
	"github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/secret-gift-api/internal/infra/bootstrap"
	"github.com/gin-gonic/gin"
)

func configureGiftsRoutes(r *gin.RouterGroup, d *bootstrap.Dependencies) {
	gifts := r.Group("gifts")
	{
		gifts.POST("", middlewares.ValidateRequest("gifts/create_gift"), d.CreateGiftController.Handle)
	}
}
