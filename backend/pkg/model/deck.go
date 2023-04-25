package model

import (
	"time"
)

type Deck struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	GroupID   uint
	Group     Group
}
