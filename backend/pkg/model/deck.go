package model

import (
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type DeckType string

const (
	Public  DeckType = "public"
	Private DeckType = "private"
)

type Deck struct {
	ID           string   `gorm:"primaryKey"`
	Name         string   `gorm:"not null"`
	DeckType     DeckType `gorm:"not null"`
	CreatedAt    time.Time
	GroupID      string `gorm:"not null"`
	Group        Group
	Cards        []Card            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserBindings []UserCardBinding `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (d *Deck) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Deck) {
		id := helper.GenerateID('D')
		return id, &Deck{ID: id.GetStringRepresentation()}
	})
	d.ID = newID.GetStringRepresentation()
	return
}
