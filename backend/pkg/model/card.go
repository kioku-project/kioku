package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Card struct {
	ID        uint `gorm:"primaryKey"`
	DeckID    uint `gorm:"not null"`
	Deck      Deck
	PublicID  string `gorm:"unique;not null"`
	Frontside string `gorm:"not null"`
	Backside  string `gorm:"not null"`
}

func (c *Card) BeforeCreate(db *gorm.DB) (err error) {
	newPublicID, err := helper.FindFreePublicID(db, 10, func() (helper.PublicID, *Card) {
		id := helper.GeneratePublicID('C')
		return id, &Card{PublicID: id.GetStringRepresentation()}
	})
	c.PublicID = newPublicID.GetStringRepresentation()
	return
}
