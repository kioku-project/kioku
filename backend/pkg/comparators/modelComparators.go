package comparators

import "github.com/kioku-project/kioku/pkg/model"

func CardModelDateComparator(a, b model.Card) int {
	return a.CreatedAt.Compare(b.CreatedAt)
}

func DeckModelDateComparator(a, b model.Deck) int {
	return a.CreatedAt.Compare(b.CreatedAt)
}

func GroupModelDateComparator(a, b model.Group) int {
	return a.CreatedAt.Compare(b.CreatedAt)
}
