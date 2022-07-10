package model

import (
	"learn-rest-api/cmd/app/model/base"
)

type UserRole struct {
	base.BaseModel
	ID 				uint 	`json:"id" gorm:"primary_key; autoIncrement; not null"`
	UserID			uint	`json:"user_id" gorm:"not null"`
	Role			string	`json:"role" gorm:"not null"`
	User			User	`json:"user" gorm:"foreignkey:UserID"`
}

func (UserRole) TableName() string {
	return "M_USER_ROLE"
}