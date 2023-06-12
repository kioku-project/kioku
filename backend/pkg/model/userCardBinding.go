package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type UserCardBinding struct {
	ID      string
	UserID  string
	CardID  string
	Card    Card
	DeckID  string `gorm:"not null"`
	Deck    Deck
	Type    uint8  `gorm:"not null"`
	Due     int64  `gorm:"not null"`
	LastIvl uint32 `gorm:"not null"`
	Factor  uint32 `gorm:"not null"`
}

func (a *UserCardBinding) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *UserCardBinding) {
		id := helper.GenerateID('S')
		return id, &UserCardBinding{ID: id.GetStringRepresentation()}
	})
	a.ID = newID.GetStringRepresentation()
	return
}
