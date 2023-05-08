package model

type Card struct {
	ID        uint `gorm:"primaryKey"`
	DeckID    uint
	Deck      Deck
	Frontside string
	Backside  string
}
