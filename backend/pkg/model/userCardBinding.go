package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type UserCardBinding struct {
	ID           string `gorm:"primaryKey"`
	UserID       string `gorm:"not null"`
	User         User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CardID       string `gorm:"not null"`
	Card         Card   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeckID       string `gorm:"not null"`
	Deck         Deck   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type         uint8  `gorm:"not null"`
	Due          int64  `gorm:"not null"`
	LastInterval uint32 `gorm:"not null"`
	Factor       uint32 `gorm:"not null"`
}

func (a *UserCardBinding) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *UserCardBinding) {
		id := helper.GenerateID('B')
		return id, &UserCardBinding{ID: id.GetStringRepresentation()}
	})
	a.ID = newID.GetStringRepresentation()
	return
}
