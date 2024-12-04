package repository

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type LimitRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Limits
	FindByID(db *gorm.DB, id *uint) domain.Limit
	Update(db *gorm.DB, limit *domain.Limit) *domain.Limit
}
