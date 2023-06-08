package store

import "github.com/kioku-project/kioku/pkg/model"

type UserStore interface {
	RegisterNewUser(newUser *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(userID string) (*model.User, error)
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
	FindCardSidesByCardID(cardID string) ([]model.CardSide, error)
	FindCardSideByID(cardSideID string) (*model.CardSide, error)
	FindLastCardSideOfCardByID(cardID string) (*model.CardSide, error)
	CreateCardSide(newCardSide *model.CardSide) error
	ModifyCardSide(cardSide *model.CardSide) error
	DeleteCardSide(cardSide *model.CardSide) error
	DeleteCardSidesOfCardByID(cardID string) error
}

type CollaborationStore interface {
	FindGroupsByUserID(userID string) ([]model.Group, error)
	FindGroupByID(groupID string) (*model.Group, error)
	CreateNewGroupWithAdmin(adminUserID string, newGroup *model.Group) error
	AddNewMemberToGroup(userID string, groupID string) error
	ModifyGroup(group *model.Group) error
	DeleteGroup(group *model.Group) error
	GetGroupUserRole(userID string, groupID string) (model.RoleType, error)
	GetGroupMemberRoles(groupID string) ([]model.GroupUserRole, error)
	CreateNewGroupAdmission(newAdmission *model.GroupAdmission) error
	FindGroupRequestsByGroupID(groupID string) ([]model.GroupAdmission, error)
	FindGroupInvitationsByUserID(userID string) ([]model.GroupAdmission, error)
	FindGroupAdmissionByUserAndGroupID(userID string, groupID string) (*model.GroupAdmission, error)
	FindGroupAdmissionByID(admissionID string) (*model.GroupAdmission, error)
	DeleteGroupAdmission(admission *model.GroupAdmission) error
}
