package store

import "github.com/kioku-project/kioku/pkg/model"

type UserStore interface {
	RegisterNewUser(newUser *model.User) error
	ModifyUser(user *model.User) error
	DeleteUser(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(userID string) (*model.User, error)
}

type CardDeckStore interface {
	FindDecksByGroupID(groupID string) ([]model.Deck, error)
	FindDeckCards(deckID string) ([]*model.Card, error)
	FindPublicDecksByGroupID(groupID string) ([]model.Deck, error)
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
	AddInvitedUserToGroup(userID string, groupID string) error
	AddRequestingUserToGroup(userID string, groupID string) error
	PromoteUserToFullGroupMember(userID string, groupID string) error
	ModifyUserRole(userID string, groupID string, role model.RoleType) error
	RemoveUserFromGroup(userID string, groupID string) error
	ModifyGroup(group *model.Group) error
	DeleteGroup(group *model.Group) error
	FindGroupUserRole(userID string, groupID string) (model.RoleType, error)
	FindGroupMemberRoles(groupID string) ([]model.GroupUserRole, error)
	FindGroupAdmins(groupID string) ([]model.GroupUserRole, error)
	FindGroupRequestsByGroupID(groupID string) ([]model.GroupUserRole, error)
	FindGroupInvitationsByUserID(userID string) ([]model.GroupUserRole, error)
	FindGroupInvitationsByGroupID(groupID string) ([]model.GroupUserRole, error)
}

type SrsStore interface {
	CreateRevlog(newRev *model.Revlog) error
	FindCardBinding(userID string, cardID string) (*model.UserCardBinding, error)
	FindUserDeckCards(userID string, deckID string) ([]*model.UserCardBinding, error)
	FindUserCards(userID string) ([]*model.UserCardBinding, error)
	CreateUserCard(newCard *model.UserCardBinding) error
	ModifyUserCard(card *model.UserCardBinding) error
}
