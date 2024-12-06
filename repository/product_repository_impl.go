package repository

import (
	"time"

	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Products {
	products := domain.Products{}
	currentDate := time.Now().Format("2006-01-02")
	tx := db.Model(&domain.Product{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Preload("ProductPrice", "start_date <= ? AND end_date >= ?", currentDate, currentDate).Preload("Company").Find(&products).Error
	helper.PanicIfError(err)

	return products
}

func (repository *ProductRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.Product {
	var product domain.Product
	currentDate := time.Now().Format("2006-01-02")

	err := db.Preload("ProductPrice", "start_date <= ? AND end_date >= ?", currentDate, currentDate).
		First(&product, id).Error
	helper.PanicIfError(err)
	return product
}
