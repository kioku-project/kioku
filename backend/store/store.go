package store

import (
	"context"

	"github.com/kioku-project/kioku/pkg/model"
)

type UserStore interface {
	RegisterNewUser(ctx context.Context, user *model.User) error
	ModifyUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.User) error
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindUserByID(ctx context.Context, userID string) (*model.User, error)
}

type CardDeckStore interface {
	FindDecksByGroupID(ctx context.Context, groupID string, userID string) ([]model.Deck, error)
	FindDeckCards(ctx context.Context, deckID string) ([]*model.Card, error)
	FindPublicDecksByGroupID(ctx context.Context, groupID string) ([]model.Deck, error)
	FindDeckByID(ctx context.Context, deckID string, userID string) (*model.Deck, error)
	CreateDeck(ctx context.Context, newDeck *model.Deck) error
	ModifyDeck(ctx context.Context, deck *model.Deck) error
	DeleteDeck(ctx context.Context, deck *model.Deck) error
	FindCardByID(ctx context.Context, cardID string) (*model.Card, error)
	CreateCard(ctx context.Context, newCard *model.Card) error
	ModifyCard(ctx context.Context, card *model.Card) error
	DeleteCard(ctx context.Context, card *model.Card) error
	FindCardSidesByCardID(ctx context.Context, cardID string) ([]model.CardSide, error)
	FindCardSideByID(ctx context.Context, cardSideID string) (*model.CardSide, error)
	FindLastCardSideOfCardByID(ctx context.Context, cardID string) (*model.CardSide, error)
	CreateCardSide(ctx context.Context, newCardSide *model.CardSide) error
	ModifyCardSide(ctx context.Context, cardSide *model.CardSide) error
	DeleteCardSide(ctx context.Context, cardSide *model.CardSide) error
	DeleteCardSidesOfCardByID(ctx context.Context, cardID string) error
	FindFavoriteDecks(ctx context.Context, userID string) ([]model.Deck, error)
	AddFavoriteDeck(ctx context.Context, userID string, deckID string) error
	DeleteFavoriteDeck(ctx context.Context, userID string, deckID string) error
	FindActiveDecks(ctx context.Context, userID string) ([]model.Deck, error)
	AddActiveDeck(ctx context.Context, userID string, deckID string) error
	DeleteActiveDeck(ctx context.Context, userID string, deckID string) error
}

type CollaborationStore interface {
	FindGroupsByUserID(ctx context.Context, userID string) ([]model.Group, error)
	FindGroupByID(ctx context.Context, groupID string) (*model.Group, error)
	CreateNewGroupWithAdmin(ctx context.Context, adminUserID string, newGroup *model.Group) error
	AddNewMemberToGroup(ctx context.Context, userID string, groupID string) error
	AddInvitedUserToGroup(ctx context.Context, userID string, groupID string) error
	AddRequestingUserToGroup(ctx context.Context, userID string, groupID string) error
	PromoteUserToFullGroupMember(ctx context.Context, userID string, groupID string) error
	ModifyUserRole(ctx context.Context, userID string, groupID string, role model.RoleType) error
	RemoveUserFromGroup(ctx context.Context, userID string, groupID string) error
	ModifyGroup(ctx context.Context, group *model.Group) error
	DeleteGroup(ctx context.Context, group *model.Group) error
	FindGroupUserRole(ctx context.Context, userID string, groupID string) (model.RoleType, error)
	FindGroupMemberRoles(ctx context.Context, groupID string) ([]model.GroupUserRole, error)
	FindGroupAdmins(ctx context.Context, groupID string) ([]model.GroupUserRole, error)
	FindGroupRequestsByGroupID(ctx context.Context, groupID string) ([]model.GroupUserRole, error)
	FindGroupInvitationsByUserID(ctx context.Context, userID string) ([]model.GroupUserRole, error)
	FindGroupInvitationsByGroupID(ctx context.Context, groupID string) ([]model.GroupUserRole, error)
}

type SrsStore interface {
	CreateRevlog(ctx context.Context, newRev *model.Revlog) error
	FindCardBinding(ctx context.Context, userID string, cardID string) (*model.UserCardBinding, error)
	FindUserDeckCards(ctx context.Context, userID string, deckID string) ([]*model.UserCardBinding, error)
	FindUserCards(ctx context.Context, userID string) ([]*model.UserCardBinding, error)
	CreateUserCard(ctx context.Context, newCard *model.UserCardBinding) error
	ModifyUserCard(ctx context.Context, card *model.UserCardBinding) error
}

type NotificationsStore interface {
	FindAllPushSubscriptions(ctx context.Context) ([]*model.PushSubscription, error)
	CreatePushSubscription(ctx context.Context, newSubscription *model.PushSubscription) error
	DeletePushSubscription(ctx context.Context, subscription *model.PushSubscription) error
	FindPushSubscriptionByID(ctx context.Context, subscriptionID string) (*model.PushSubscription, error)
}
