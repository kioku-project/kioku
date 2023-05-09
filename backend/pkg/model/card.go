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

func (c *Card) BeforeCreate(db *gorm.DB) error {
	var card Card
	isUnique := false
	for !isUnique {
		randomPublicID := helper.GeneratePublicID()
		err := db.Where(Card{PublicID: helper.ConvertRandomIDToModelID('C', randomPublicID)}).First(&card).Error
		if err != nil {
			isUnique = true
			c.PublicID = helper.ConvertRandomIDToModelID('C', randomPublicID)
		}
	}
	return nil
}
