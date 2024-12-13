package service

import (
	"fmt"
	"time"

	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/AsrofunNiam/technical-test-credit-plus/exception"
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

	//  Validate request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	channel := make(chan domain.Product)
	defer close(channel)

	// Validate user role
	if auth.Role != "customer" {
		err := &exception.ErrorSendToResponse{Err: "Only customers can create loans"}
		helper.PanicIfError(err)
	}

	// Find product
	// product := service.ProductRepository.FindByID(tx, &request.ProductID)

	// implement chanel to find product
	go func() {
		product := service.ProductRepository.FindByID(tx, &request.ProductID)
		channel <- product
		fmt.Println("done query product")

	}()

	product := <-channel

	// Find limit
	limit := service.LimitRepository.FindByID(tx, &auth.ID, &request.Period)

	// Validate limit existence
	if limit.ID == 0 {
		err := &exception.ErrorSendToResponse{Err: "Limit not found"}
		helper.PanicIfError(err)
	}

	instalmentAmount := product.ProductPrice.Price + request.AdminFee

	// Validate limit against instalment amount
	if limit.LimitAmount < instalmentAmount {
		err := &exception.ErrorSendToResponse{Err: "Insufficient limit for the requested product"}
		helper.PanicIfError(err)
	}

	// Default interest on year
	interestYearValue := 10.0

	// Calculate interest (metode flat)
	interestAmount := instalmentAmount * (interestYearValue / 100) * (float64(request.Period) / 12)

	transaction := &domain.Transaction{
		// Required Fields
		CreatedByID: auth.ID,
		UserID:      auth.ID,

		// Fields for Loan
		OnTheRoad:        product.ProductPrice.Price,
		AdminFee:         request.AdminFee,
		InstalmentAmount: instalmentAmount,
		InterestAmount:   interestAmount,
		Period:           request.Period,
		ProductID:        request.ProductID,
		ProductName:      product.Name,
		TransactionType:  request.TransactionType,
		NumberContract:   fmt.Sprintf("%d-%d-%d", auth.ID, product.ID, time.Now().Unix()),
	}

	transaction = service.TransactionRepository.Create(tx, transaction)

	// Update limit
	limit.LimitAmount -= instalmentAmount
	service.LimitRepository.Update(tx, &limit)

	return transaction.ToTransactionResponse()
}
