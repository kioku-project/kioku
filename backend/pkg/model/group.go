package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Group struct {
	ID        string          `gorm:"primaryKey"`
	Name      string          `gorm:"not null"`
	IsDefault bool            `gorm:"not null"`
	Users     []User          `gorm:"many2many:group_user_roles;"`
	Decks     []Deck          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Roles     []GroupUserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (g *Group) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Group) {
		id := helper.GenerateID('G')
		return id, &Group{ID: id.GetStringRepresentation()}
	})
	g.ID = newID.GetStringRepresentation()
	return
}
