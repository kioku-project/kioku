package store

import "github.com/kioku-project/kioku/pkg/model"

type UserStore interface {
	FindUserByEmail(email string) (*model.User, error)
	RegisterNewUser(newUser *model.User) error
}

type CardDeckStore interface {
	CreateDeck(newDeck *model.Deck) error
	FindDeckByName(deckName string) (*model.Deck, error)
	CreateCard(newCard *model.Card) error
}

type CollaborationStore interface {
	CreateNewGroupWithAdmin(adminUserID uint, newGroup *model.Group) error
	FindGroupByName(groupName string) (*model.Group, error)
	GetGroupUserRole(userID uint, groupID uint) (*model.RoleType, error)
}
