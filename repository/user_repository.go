package repository

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUserName(db *gorm.DB, userName *string) domain.User
}
