package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type User struct {
	ID             string  `gorm:"primaryKey;" json:"userID,omitempty"`
	Name           string  `gorm:"not null;" json:"userName,omitempty"`
	Email          string  `gorm:"unique;not null;" json:"userEmail,omitempty"`
	Password       string  `gorm:"not null;" json:"userPassword,omitempty"`
	FavoritesDecks []*Deck `gorm:"many2many:user_favorite_decks"`
	ActiveDecks    []*Deck `gorm:"many2many:user_active_decks"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *User) {
		id := helper.GenerateID('U')
		return id, &User{ID: id.GetStringRepresentation()}
	})
	u.ID = newID.GetStringRepresentation()
	return
}
