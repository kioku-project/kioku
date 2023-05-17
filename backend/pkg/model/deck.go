package model

import (
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Deck struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	GroupID   string    `gorm:"not null"`
	Group     Group     `gorm:"foreignKey:GroupID"`
	Cards     []Card    `gorm:"foreignKey:DeckID"`
}

func (d *Deck) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Deck) {
		id := helper.GenerateID('D')
		return id, &Deck{ID: id.GetStringRepresentation()}
	})
	d.ID = newID.GetStringRepresentation()
	return
}
