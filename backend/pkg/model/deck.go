package model

import (
	"time"
)

type Deck struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	GroupID   uint      `gorm:"not null"`
	Group     Group
}
