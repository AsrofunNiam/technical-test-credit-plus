package repository

import (
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
	tx := db.Model(&domain.Product{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&products).Error
	helper.PanicIfError(err)

	return products
}

func (repository *ProductRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.Product {
	var product domain.Product
	err := db.First(&product, id).Error
	helper.PanicIfError(err)
	return product
}
