package service

import (
	"net/http"

	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"github.com/AsrofunNiam/technical-test-credit-plus/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	ProductRepository     repository.ProductRepository
	LimitRepository       repository.LimitRepository
	DB                    *gorm.DB
	Validate              *validator.Validate
}

func NewTransactionService(
	transactionRepository repository.TransactionRepository,
	productRepository repository.ProductRepository,
	limitRepository repository.LimitRepository,
	db *gorm.DB,
	validate *validator.Validate,
) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		ProductRepository:     productRepository,
		LimitRepository:       limitRepository,
		DB:                    db,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.TransactionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	transaction := service.TransactionRepository.FindAll(tx, filters)
	return transaction.ToTransactionResponses()
}

func (service *TransactionServiceImpl) CreateLoanSubmission(auth *auth.AccessDetails, request *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//  check for limit
	limit := service.LimitRepository.FindByID(tx, &auth.ID)

	//  validate limit
	if limit.ID == 0 {
		helper.SendErrorResponse(c, http.StatusNotFound, "Limit not found")
		// return web.TransactionResponse{}
	}

	//  validate limit transaction
	if limit.LimitAmount < request.InstalmentAmount {
		helper.SendErrorResponse(c, http.StatusBadRequest, "Insufficient limit")
		// return web.TransactionResponse{}
	}

	transaction := &domain.Transaction{
		// Required Fields
		CreatedByID: auth.ID,
		UserID:      request.UserID,

		// Optional Fields for Loan
		OnTheRoad:        request.OnTheRoad,
		AdminFee:         request.AdminFee,
		InstalmentAmount: request.InstalmentAmount,
		InterestAmount:   request.InterestAmount,
		Period:           request.Period,
		ProductID:        request.ProductID,
		TransactionType:  request.TransactionType,
	}

	transaction = service.TransactionRepository.Create(tx, transaction)

	// update limit
	limit.LimitAmount -= request.InstalmentAmount
	service.LimitRepository.Update(tx, &limit)

	return transaction.ToTransactionResponse()
}
