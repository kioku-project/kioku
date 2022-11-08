package card

import "github.com/kioku-project/kioku/pkg/deck"

type Card struct {
	ID     uint
	DeckID uint
	Deck   deck.Deck
}
