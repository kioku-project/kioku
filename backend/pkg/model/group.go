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

func (g *Group) BeforeCreate(db *gorm.DB) error {
	var group Group
	isUnique := false
	for !isUnique {
		randomPublicID := helper.GeneratePublicID()
		err := db.Where(Group{PublicID: helper.ConvertRandomIDToModelID('G', randomPublicID)}).First(&group).Error
		if err != nil {
			isUnique = true
			g.PublicID = helper.ConvertRandomIDToModelID('G', randomPublicID)
		}
	}
	return nil
}
