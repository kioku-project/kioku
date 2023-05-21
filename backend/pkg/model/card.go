package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Card struct {
	ID        string `gorm:"primaryKey"`
	DeckID    string `gorm:"not null"`
	Deck      Deck
	Frontside string `gorm:"not null"`
	Backside  string `gorm:"not null"`
}

func (c *Card) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Card) {
		id := helper.GenerateID('C')
		return id, &Card{ID: id.GetStringRepresentation()}
	})
	c.ID = newID.GetStringRepresentation()
	return
}
