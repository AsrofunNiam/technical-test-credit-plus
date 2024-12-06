package service

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"github.com/gin-gonic/gin"
)

type TransactionService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.TransactionResponse
	CreateLoanSubmission(auth *auth.AccessDetails, request *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse
}
