package service

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"github.com/AsrofunNiam/technical-test-credit-plus/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewProductService(
	product repository.ProductRepository,
	db *gorm.DB,
	validate *validator.Validate,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository: product,
		DB:                db,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse {
	tx := service.DB
	err := tx.Error
	helper.PanicIfError(err)

	products := service.ProductRepository.FindAll(tx, filters)
	return products.ToProductResponses()
}

func (service *ProductServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.ProductResponse {
	tx := service.DB
	err := tx.Error
	helper.PanicIfError(err)

	product := service.ProductRepository.FindByID(tx, id)
	return product.ToProductResponse()
}
