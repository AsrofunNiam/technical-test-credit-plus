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

func TransactionRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {
	Transactions := service.NewTransactionService(
		repository.NewTransactionRepository(),
		repository.NewProductRepository(),
		repository.NewLimitRepository(),
		db,
		validate,
	)
	productController := controller.NewTransactionController(Transactions)
	router.GET("/transactions", auth.Auth(productController.FindAll, []string{}))
	router.POST("/transactions/loan-submission", auth.Auth(productController.CreateLoanSubmission, []string{}))
}
