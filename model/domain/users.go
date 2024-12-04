package domain

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"gorm.io/gorm"
)

type Users []User
type User struct {
	// Required Fields
	gorm.Model
	CreatedByID uint  `gorm:""`
	UpdatedByID uint  `gorm:""`
	DeletedByID *uint `gorm:""`

	// Fields
	FullName      string `gorm:"size:200;uniqueIndex:idx_users"`
	LegalName     string `gorm:"size:200"`
	PlaceOfBirth  string `gorm:"size:200"`
	DateOfBirth   string `gorm:"size:200"`
	Salary        float64
	IdentityImage string `gorm:"size:200"`
	FaceImage     string `gorm:"size:200"`
	Password      string `gorm:"size:100"`
}

func (user *User) ToUserResponse() web.UserResponse {
	return web.UserResponse{
		// Required Fields
		ID: user.ID,
		// Fields
		FullName:  user.FullName,
		LegalName: user.LegalName,
	}
}

func (users Users) ToUserResponses() []web.UserResponse {
	userResponses := []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, user.ToUserResponse())
	}
	return userResponses
}
