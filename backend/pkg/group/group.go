package group

import (
	"github.com/kioku-project/kioku/pkg/deck"
	"github.com/kioku-project/kioku/pkg/user"
)

type Group struct {
	ID    uint
	Name  string
	Users []user.User `gorm:"many2many:group_users"`
	Decks []deck.Deck
}
