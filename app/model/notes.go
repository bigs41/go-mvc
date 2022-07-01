package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"` // Explicitly specify the type to be uuid
	Title    string
	SubTitle string
	Text     string
}

func (u *Note) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
