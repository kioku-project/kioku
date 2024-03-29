package model

import (
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type DeckType string

const (
	PublicDeckType  DeckType = "public"
	PrivateDeckType DeckType = "private"
)

type Deck struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	DeckType    DeckType `gorm:"not null"`
	CreatedAt   time.Time
	GroupID     string `gorm:"not null"`
	Group       Group
	Cards       []Card `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsFavorite  bool   `gorm:"-"`
	IsActive    bool   `gorm:"-"`
}

func (d *Deck) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Deck) {
		id := helper.GenerateID('D')
		return id, &Deck{ID: id.GetStringRepresentation()}
	})
	d.ID = newID.GetStringRepresentation()
	return
}
