package model

import (
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Deck struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	PublicID  string    `gorm:"unique;not null"`
	GroupID   uint      `gorm:"not null"`
	Group     Group
}

func (d *Deck) BeforeCreate(db *gorm.DB) error {
	var deck Deck
	isUnique := false
	for !isUnique {
		randomPublicID := helper.GeneratePublicID()
		err := db.Where(Deck{PublicID: helper.ConvertRandomIDToModelID('D', randomPublicID)}).First(&deck).Error
		if err != nil {
			isUnique = true
			d.PublicID = helper.ConvertRandomIDToModelID('D', randomPublicID)
		}
	}
	return nil
}
