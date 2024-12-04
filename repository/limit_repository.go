package repository

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type LimitRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Limits
	FindByID(db *gorm.DB, userID *uint, tenor *int) domain.Limit
	Update(db *gorm.DB, limit *domain.Limit) *domain.Limit
}
