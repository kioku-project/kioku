package store

import "github.com/kioku-project/kioku/pkg/model"

type UserStore interface {
	RegisterNewUser(newUser *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type CardDeckStore interface {
	FindDecksByGroupID(groupID string) ([]model.Deck, error)
	FindDeckByID(deckID string) (*model.Deck, error)
	CreateDeck(newDeck *model.Deck) error
	ModifyDeck(deck *model.Deck) error
	DeleteDeck(deck *model.Deck) error
	FindCardByID(cardID string) (*model.Card, error)
	CreateCard(newCard *model.Card) error
	ModifyCard(card *model.Card) error
	DeleteCard(card *model.Card) error
}

type CollaborationStore interface {
	FindGroupsByUserID(userID string) ([]model.Group, error)
	FindGroupByID(groupID string) (*model.Group, error)
	CreateNewGroupWithAdmin(adminUserID string, newGroup *model.Group) error
	ModifyGroup(group *model.Group) error
	DeleteGroup(group *model.Group) error
	GetGroupUserRole(userID string, groupID string) (model.RoleType, error)
}
