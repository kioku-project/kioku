package model

type AlgorithmType string

const (
	AlgoDynamicSRS AlgorithmType = "Dynamic-Spaced-Repetition"
	AlgoLinearSRS                = "Linear-Spaced-Repetition"
	AlgoAISRS                    = "AI-Spaced-Repetition"
)

type UserActiveDecks struct {
	Algorithm AlgorithmType `gorm:"not null"`
	UserID    string        `gorm:"not null;primaryKey"`
	User      User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeckID    string        `gorm:"not null;primaryKey"`
	Deck      Deck          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserFavoriteDecks struct {
	UserID string `gorm:"not null;primaryKey"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeckID string `gorm:"not null;primaryKey"`
	Deck   Deck   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
