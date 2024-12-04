package repository

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindByUserName(db *gorm.DB, userName *string) domain.User {
	var user domain.User
	err := db.Where("user_name = ?", userName).First(&user).Error
	helper.PanicIfError(err)
	return user
}
