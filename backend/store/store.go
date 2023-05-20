package store

import "github.com/kioku-project/kioku/pkg/model"

type UserStore interface {
	FindUserByEmail(email string) (*model.User, error)
	RegisterNewUser(newUser *model.User) error
}

type CardDeckStore interface {
	CreateDeck(newDeck *model.Deck) error
	FindDeckByID(deckID string) (*model.Deck, error)
	CreateCard(newCard *model.Card) error
	FindDecksByGroupID(groupID string) ([]model.Deck, error)
}

type CollaborationStore interface {
	CreateNewGroupWithAdmin(adminUserID string, newGroup *model.Group) error
	FindGroupByID(groupID string) (*model.Group, error)
	GetGroupUserRole(userID string, groupID string) (model.RoleType, error)
	FindGroupsByUserID(userID string) ([]model.Group, error)
}
