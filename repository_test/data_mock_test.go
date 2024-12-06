package repository_test

import (
	"time"

	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

func MockProductFind() *domain.Product {
	productPrice := []domain.ProductPrice{
		{
			ProductID: 1,
			Price:     15000000.00,
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
	}

	return &domain.Product{
		Model:        gorm.Model{ID: 1},
		Name:         "Honda Beat",
		ProductPrice: productPrice[0],
	}
}

func MockTransactionCreate() *domain.Transaction {
	transaction := &domain.Transaction{
		CreatedByID:      1,
		UpdatedByID:      1,
		DeletedByID:      0,
		UserID:           2,
		NumberContract:   "12345",
		OnTheRoad:        500000,
		AdminFee:         10000,
		InstalmentAmount: 2000000,
		InterestAmount:   100000,
		ProductID:        1,
		ProductName:      "Honda Beat",
		Period:           12,
		TransactionType:  "purchase",
	}

	return transaction
}

func MockLimitUpdate() *domain.Limit {
	return &domain.Limit{
		UserID:      1,
		Tenor:       12,
		LimitAmount: 5000000,
		Available:   true,
	}
}
