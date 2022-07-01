package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"` // Explicitly specify the type to be uuid
	Name      string    `json:"name" xml:"name" form:"name"`
	CompanyID uuid.UUID `json:"company_id" xml:"company_id" form:"company_id"`
	Company   *Company
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
