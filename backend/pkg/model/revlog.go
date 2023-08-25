package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Revlog struct {
	ID     string `gorm:"primaryKey"`
	CardID string `gorm:"not null"`
	Card   Card
	UserID string `gorm:"not null"`
        User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date   int64 `gorm:"not null"`
	Rating int64 `gorm:"not null"`
}

func (r *Revlog) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Revlog) {
		id := helper.GenerateID('R')
		return id, &Revlog{ID: id.GetStringRepresentation()}
	})
	r.ID = newID.GetStringRepresentation()
	return
}
