package repository

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(db *gorm.DB, identity *string) domain.User
}
