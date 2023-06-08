package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type CardSide struct {
	ID                 string `gorm:"primaryKey"`
	CardID             string `gorm:"not null"`
	Card               Card
	Header             string `gorm:"not null"`
	Description        string `gorm:"not null"`
	PreviousCardSideID string
	NextCardSideID     string
}

func (c *CardSide) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *CardSide) {
		id := helper.GenerateID('S')
		return id, &CardSide{ID: id.GetStringRepresentation()}
	})
	c.ID = newID.GetStringRepresentation()
	return
}
