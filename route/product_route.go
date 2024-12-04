package route

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/AsrofunNiam/technical-test-credit-plus/controller"
	"github.com/AsrofunNiam/technical-test-credit-plus/repository"
	"github.com/AsrofunNiam/technical-test-credit-plus/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {
	Products := service.NewProductService(
		repository.NewProductRepository(),
		db,
		validate,
	)
	productController := controller.NewProductController(Products)
	router.GET("/products", auth.Auth(productController.FindAll, []string{}))
	router.GET("/products/:id", auth.Auth(productController.FindByID, []string{}))
}
