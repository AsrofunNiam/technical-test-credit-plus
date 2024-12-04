package controller

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	FindByID(context *gin.Context, auth *auth.AccessDetails)
}
