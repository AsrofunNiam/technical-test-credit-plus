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
	limits := domain.Limits{}
	tx := db.Model(&domain.Limit{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&limits).Error
	helper.PanicIfError(err)

	return limits
}

func (repository *LimitRepositoryImpl) FindByID(db *gorm.DB, userID *uint, tenor *int) domain.Limit {
	var limit domain.Limit
	err := db.Where("user_id = ? AND tenor = ?", userID, tenor).First(&limit).Error
	helper.PanicIfError(err)
	return limit
}

func (repository *LimitRepositoryImpl) Update(db *gorm.DB, limit *domain.Limit) *domain.Limit {
	err := db.Updates(&limit).First(&limit).Error
	helper.PanicIfError(err)
	return limit
}
