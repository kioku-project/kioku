package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey;" json:"userID"`
	Name     string `gorm:"not null;" json:"userName"`
	Email    string `gorm:"unique;not null;" json:"userEmail"`
	Password string `gorm:"not null;" json:"userPassword"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *User) {
		id := helper.GenerateID('U')
		return id, &User{ID: id.GetStringRepresentation()}
	})
	u.ID = newID.GetStringRepresentation()
	return
}
