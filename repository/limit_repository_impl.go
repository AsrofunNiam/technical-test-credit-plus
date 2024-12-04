package repository

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type LimitRepositoryImpl struct {
}

func NewLimitRepository() LimitRepository {
	return &LimitRepositoryImpl{}
}

func (repository *LimitRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Limits {
	Limits := domain.Limits{}
	tx := db.Model(&domain.Limit{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&Limits).Error
	helper.PanicIfError(err)

	return Limits
}

func (repository *LimitRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.Limit {
	var Limit domain.Limit
	err := db.First(&Limit, id).Error
	helper.PanicIfError(err)
	return Limit
}

func (repository *LimitRepositoryImpl) Update(db *gorm.DB, limit *domain.Limit) *domain.Limit {
	err := db.Updates(&limit).First(&limit).Error
	helper.PanicIfError(err)
	return limit
}
