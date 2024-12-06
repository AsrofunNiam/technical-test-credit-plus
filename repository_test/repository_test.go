package repository_test

import (
	"testing"

	"github.com/AsrofunNiam/technical-test-credit-plus/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Preload(column string, query interface{}, args ...interface{}) *gorm.DB {
	args = append(args, args...)
	m.Called(column, query, args)
	return &gorm.DB{}
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.Called(dest, conds)
	return &gorm.DB{}
}

func TestFindData(t *testing.T) {
	// Setup
	tx := Connection()
	repositoryProduct := repository.ProductRepositoryImpl{}

	// Test Product
	product := MockProductFind()
	result := repositoryProduct.FindByID(tx, &product.ID)

	// Assert result Product
	assert.NotNil(t, result)
	assert.Equal(t, product.Name, result.Name)
	assert.Equal(t, product.ProductPrice.Price, result.ProductPrice.Price)

}

func TestCreate(t *testing.T) {
	// Setup
	tx := Connection()
	tx = tx.Begin()
	repositoryTransaction := repository.TransactionRepositoryImpl{}
	if tx.Error != nil {
		t.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	// Test Transaction
	transaction := MockTransactionCreate()

	result := repositoryTransaction.Create(tx, transaction)

	// Assert result Transaction
	assert.NotNil(t, result)
	assert.Equal(t, transaction.UserID, result.UserID)
	assert.Equal(t, transaction.NumberContract, result.NumberContract)
	assert.Equal(t, transaction.OnTheRoad, result.OnTheRoad)
	assert.Equal(t, transaction.AdminFee, result.AdminFee)
	assert.Equal(t, transaction.InstalmentAmount, result.InstalmentAmount)
	assert.Equal(t, transaction.InterestAmount, result.InterestAmount)
	assert.Equal(t, transaction.ProductID, result.ProductID)
	assert.Equal(t, transaction.ProductName, result.ProductName)
	assert.Equal(t, transaction.Period, result.Period)
	assert.Equal(t, transaction.TransactionType, result.TransactionType)

	if assert.NoError(t, tx.Error) {
		tx.Commit()
	} else {
		tx.Rollback()
		t.Fatalf("test failed, changes rolled back")
	}

}

func TestUpdate(t *testing.T) {
	// Setup
	tx := Connection()
	tx = tx.Begin()
	repositoryLimit := repository.LimitRepositoryImpl{}
	if tx.Error != nil {
		t.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	limit := MockLimitUpdate()

	limitFind := repositoryLimit.FindByID(tx, &limit.UserID, &limit.Tenor)
	limitFinal := limitFind.LimitAmount - limit.LimitAmount

	limitFind.LimitAmount -= limit.LimitAmount

	updatedLimit := repositoryLimit.Update(tx, &limitFind)

	// Assert result Limit
	assert.NotNil(t, updatedLimit)
	assert.Equal(t, limit.UserID, updatedLimit.UserID)
	assert.Equal(t, limitFinal, updatedLimit.LimitAmount)
	assert.Equal(t, limit.Tenor, updatedLimit.Tenor)

	if assert.NoError(t, tx.Error) {
		tx.Commit()
	} else {
		tx.Rollback()
		t.Fatalf("test failed, changes rolled back")
	}
}
