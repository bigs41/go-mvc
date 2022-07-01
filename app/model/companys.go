package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primary_key;"` // Explicitly specify the type to be uuid
	Name string
}

func (u *Company) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
