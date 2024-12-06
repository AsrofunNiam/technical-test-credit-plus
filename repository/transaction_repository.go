package repository

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Transactions
	Create(db *gorm.DB, transaction *domain.Transaction) *domain.Transaction
}
