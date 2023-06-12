package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Card struct {
	ID              string `gorm:"primaryKey"`
	DeckID          string `gorm:"not null"`
	Deck            Deck
	FirstCardSideID string            `gorm:"not null"`
	CardSides       []CardSide        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserBindings    []UserCardBinding `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (c *Card) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Card) {
		id := helper.GenerateID('C')
		return id, &Card{ID: id.GetStringRepresentation()}
	})
	c.ID = newID.GetStringRepresentation()
	return
}
