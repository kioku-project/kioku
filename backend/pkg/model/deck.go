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
	newPublicID, err := helper.FindFreePublicID(db, 10, func() (helper.PublicID, *Deck) {
		id := helper.GeneratePublicID('D')
		return id, &Deck{PublicID: id.GetStringRepresentation()}
	})
	d.PublicID = newPublicID.GetStringRepresentation()
	return
}
