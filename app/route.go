package app

import (
	route "github.com/AsrofunNiam/technical-test-credit-plus/route"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {

	router := gin.New()
	router.UseRawPath = true

	route.UserRoute(router, db, validate)
	route.ProductRoute(router, db, validate)
	route.TransactionRoute(router, db, validate)

	return router
}
