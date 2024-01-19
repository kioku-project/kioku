package model

import gorm "gorm.io/gorm"
import "github.com/kioku-project/kioku/pkg/helper"

type PushNotification struct {
	Title   string              `json:"title"`
	Body    string              `json:"body"`
	Actions []map[string]string `json:"actions"`
	Vibrate []int               `json:"vibrate"`
}

type PushSubscription struct {
	ID       string `gorm:"primaryKey"`
	UserID   string `gorm:"not null"`
	User     User
	Endpoint string `gorm:"not null"`
	P256DH   string `gorm:"not null"`
	Auth     string `gorm:"not null"`
}

func (c *PushSubscription) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *PushSubscription) {
		id := helper.GenerateID('N')
		return id, &PushSubscription{ID: id.GetStringRepresentation()}
	})
	c.ID = newID.GetStringRepresentation()
	return
}
