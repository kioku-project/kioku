package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type Group struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	PublicID string `gorm:"unique;not null"`
	Users    []User `gorm:"many2many:group_user_roles;"`
}

func (g *Group) BeforeCreate(db *gorm.DB) (err error) {
	newPublicID, err := helper.FindFreePublicID(db, 10, func() (helper.PublicID, *Group) {
		id := helper.GeneratePublicID('G')
		return id, &Group{PublicID: id.GetStringRepresentation()}
	})
	g.PublicID = newPublicID.GetStringRepresentation()
	return
}
