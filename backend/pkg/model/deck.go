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
	Cards     []Card `gorm:"foreignKey:DeckID"`
}

func (d *Deck) BeforeCreate(db *gorm.DB) (err error) {
	d.PublicID, err = helper.FindFreePublicID(db, 10, helper.GeneratePublicID, 'D', func(candidate string) *Deck {
		return &Deck{PublicID: candidate}
	})
	return
}
