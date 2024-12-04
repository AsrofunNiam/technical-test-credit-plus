package controller

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	CreateLoanSubmission(context *gin.Context, auth *auth.AccessDetails)
}
