package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Group struct {
	ID    string `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Users []User `gorm:"many2many:group_user_roles;"`
}

func (g *Group) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *Group) {
		id := helper.GenerateID('G')
		return id, &Group{ID: id.GetStringRepresentation()}
	})
	g.ID = newID.GetStringRepresentation()
	return
}
