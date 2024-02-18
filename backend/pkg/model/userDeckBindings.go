package model

type AlgorithmType string

const (
	AlgoDynamicSRS AlgorithmType = "Dynamic-Spaced-Repetition"
	AlgoLinearSRS                = "Linear-Spaced-Repetition"
	AlgoStaticSRS                = "Static-Spaced-Repetition"
	AlgoAISRS                    = "AI-Spaced-Repetition"
)

type UserActiveDecks struct {
	UserID         string        `gorm:"not null;primaryKey"`
	User           User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeckID         string        `gorm:"not null;primaryKey"`
	Deck           Deck          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Algorithm      AlgorithmType `gorm:"not null"`
	NewCardsPerDay uint64        `gorm:"not null"`
}

type UserFavoriteDecks struct {
	UserID string `gorm:"not null;primaryKey"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeckID string `gorm:"not null;primaryKey"`
	Deck   Deck   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
