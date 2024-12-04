package controller

import (
	"net/http"

	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"github.com/AsrofunNiam/technical-test-credit-plus/service"
	"github.com/gin-gonic/gin"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like", "id.eq")
	transactionResponses := controller.TransactionService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(transactionResponses),
		Data:    transactionResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionControllerImpl) CreateLoanSubmission(c *gin.Context, auth *auth.AccessDetails) {
	request := &web.TransactionCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	transactionResponse := controller.TransactionService.CreateLoanSubmission(auth, request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Transaction created successfully",
		Data:    transactionResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}
