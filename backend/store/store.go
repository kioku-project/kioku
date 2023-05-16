package store

import "github.com/kioku-project/kioku/pkg/model"

type UserStore interface {
	FindUserByEmail(email string) (*model.User, error)
	RegisterNewUser(newUser *model.User) error
}

type CardDeckStore interface {
	CreateDeck(newDeck *model.Deck) error
	FindDeckByPublicID(publicID string) (*model.Deck, error)
	CreateCard(newCard *model.Card) error
	FindDecksByGroupID(groupID uint) ([]model.Deck, error)
}

type CollaborationStore interface {
	CreateNewGroupWithAdmin(adminUserID uint, newGroup *model.Group) error
	FindGroupByPublicID(publicID string) (*model.Group, error)
	GetGroupUserRole(userID uint, groupID uint) (*model.RoleType, error)
	FindGroupsByUserID(userID uint) ([]model.Group, error)
}
