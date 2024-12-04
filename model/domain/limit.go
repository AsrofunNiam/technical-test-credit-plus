package domain

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"gorm.io/gorm"
)

type Limits []Limit
type Limit struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	Tenor       int     `gorm:"not null"`
	LimitAmount float64 `gorm:"type:decimal(10,2);not null" `
	Available   bool    `gorm:"default:true"`

	// Relations
	User User `gorm:"foreignKey:UserID"`
}

func (limit *Limit) ToLimitResponse() web.LimitResponse {
	return web.LimitResponse{
		// Required Fields
		ID:          limit.ID,
		UserID:      limit.UserID,
		User:        limit.User.ToUserResponse(),
		Tenor:       limit.Tenor,
		LimitAmount: limit.LimitAmount,
		Available:   limit.Available,

		// Relations
		// CompanyResponse: limit.Company.ToCompanyResponse(),
	}
}

func (users Limits) ToLimitResponses() []web.LimitResponse {
	limitResponses := []web.LimitResponse{}
	for _, user := range users {
		limitResponses = append(limitResponses, user.ToLimitResponse())
	}
	return limitResponses
}
