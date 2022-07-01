package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"` // Explicitly specify the type to be uuid
	// add more common columns like CreatedAt
	// CreatedAt *time.Time
	// ...
}

func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
