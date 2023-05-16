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
	g.PublicID, err = helper.FindFreePublicID(db, 10, helper.GeneratePublicID, 'G', func(candidate string) *Group {
		return &Group{PublicID: candidate}
	})
	return
}
