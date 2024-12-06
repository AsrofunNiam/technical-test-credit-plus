package domain

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"gorm.io/gorm"
)

type Products []Product
type Product struct {
	gorm.Model
	CreatedByID uint `gorm:""`
	UpdatedByID uint `gorm:""`
	DeletedByID uint `gorm:""`

	// Required Fields
	Name        string `gorm:"type:varchar(255);not null"`
	Type        string `gorm:"type:text"`
	CompanyCode uint   `gorm:"not null"`
	Description string `gorm:"type:text"`
	Available   bool   `gorm:"default:true"`

	// Relations
	Company      Company      `gorm:"foreignKey:CompanyCode;references:ID"`
	ProductPrice ProductPrice `gorm:"foreignKey:ProductID;references:ID"`
}

func (product *Product) ToProductResponse() web.ProductResponse {
	return web.ProductResponse{
		// Required Fields
		ID:          product.ID,
		Name:        product.Name,
		Type:        product.Type,
		CompanyCode: product.CompanyCode,
		Description: product.Description,
		Available:   product.Available,

		// Relations
		Company:      product.Company.ToCompanyResponse(),
		ProductPrice: product.ProductPrice.ToProductPriceResponse(),
	}
}

func (users Products) ToProductResponses() []web.ProductResponse {
	productResponses := []web.ProductResponse{}
	for _, user := range users {
		productResponses = append(productResponses, user.ToProductResponse())
	}
	return productResponses
}
