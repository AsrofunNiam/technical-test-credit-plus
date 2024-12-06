package domain

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"gorm.io/gorm"
)

type Transactions []Transaction
type Transaction struct {
	gorm.Model
	CreatedByID uint `gorm:""`
	UpdatedByID uint `gorm:""`
	DeletedByID uint `gorm:""`

	// foreign key
	UserID           uint    `gorm:"not null"`
	NumberContract   string  `gorm:"type:varchar(50);not null"`
	OnTheRoad        float64 `gorm:"type:decimal(20,2);not null"`
	AdminFee         float64 `gorm:"type:decimal(20,2);not null"`
	InstalmentAmount float64 `gorm:"type:decimal(20,2);not null"`
	InterestAmount   float64 `gorm:"type:decimal(20,2);not null"`
	ProductID        uint    `gorm:"not null"`
	ProductName      string  `gorm:"type:varchar(255);not null"`
	Period           int     `gorm:"not null"`
	TransactionType  string  `gorm:"type:enum('withdrawal', 'purchase');not null"`

	// relationship
	Product Product `gorm:"foreignKey:ProductID"`
	User    User    `gorm:"foreignKey:UserID"`
}

func (transaction *Transaction) ToTransactionResponse() web.TransactionResponse {
	return web.TransactionResponse{
		// Required Fields
		ID:               transaction.ID,
		UserID:           transaction.UserID,
		NumberContract:   transaction.NumberContract,
		OnTheRoad:        transaction.OnTheRoad,
		AdminFee:         transaction.AdminFee,
		InstalmentAmount: transaction.InstalmentAmount,
		InterestAmount:   transaction.InterestAmount,
		Period:           transaction.Period,
		TransactionType:  transaction.TransactionType,

		// Relations
		Product: transaction.Product.ToProductResponse(),
		User:    transaction.User.ToUserResponse(),
	}
}

func (transactions Transactions) ToTransactionResponses() []web.TransactionResponse {
	transactionResponses := []web.TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, transaction.ToTransactionResponse())
	}
	return transactionResponses
}
