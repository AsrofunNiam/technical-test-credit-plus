package service

import (
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

	//  validate request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// validate user
	if auth.Role != "customer" {
		err := &exception.ErrorSendToResponse{Err: "Only customer can create loan"}
		helper.PanicIfError(err)
	}

	//  find limit
	limit := service.LimitRepository.FindByID(tx, &auth.ID, &request.Period)

	//  validate limit
	if limit.ID == 0 {
		err := &exception.ErrorSendToResponse{Err: "Limit not found"}
		helper.PanicIfError(err)
	}

	// Validasi apakah limit cukup untuk harga barang (OnTheRoad)
	if limit.LimitAmount < request.OnTheRoad {
		err := &exception.ErrorSendToResponse{Err: "Insufficient limit for the requested product"}
		helper.PanicIfError(err)
	}

	// Perhitungan total biaya pinjaman
	totalInterest := request.OnTheRoad * request.InterestAmount * float64(request.Period) / 100
	totalAmount := request.AdminFee + totalInterest

	// Hitung cicilan per bulan
	instalmentAmount := totalAmount / float64(request.Period)

	// Validasi apakah limit cukup untuk cicilan bulanan
	if limit.LimitAmount < totalAmount {
		err := &exception.ErrorSendToResponse{Err: "Insufficient limit for the requested product "}
		helper.PanicIfError(err)
	}

	transaction := &domain.Transaction{
		// Required Fields
		CreatedByID: auth.ID,
		UserID:      request.UserID,

		// Optional Fields for Loan
		OnTheRoad:        request.OnTheRoad,
		AdminFee:         request.AdminFee,
		InstalmentAmount: instalmentAmount,
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
