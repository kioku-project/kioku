package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type User struct {
	ID       string  `gorm:"primaryKey;" json:"id"`
	Name     string  `gorm:"not null;" json:"name"`
	Email    string  `gorm:"unique;not null;" json:"email"`
	Password string  `gorm:"not null;" json:"password"`
	Groups   []Group `gorm:"many2many:group_user_roles;" json:"groups"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *User) {
		id := helper.GenerateID('U')
		return id, &User{ID: id.GetStringRepresentation()}
	})
	u.ID = newID.GetStringRepresentation()
	return
}
