package store

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/model"
)

type GormStore struct {
	db *gorm.DB
}

type UserStoreImpl GormStore

type CardDeckStoreImpl GormStore

type CollaborationStoreImpl GormStore

type SrsStoreImpl GormStore

func NewPostgresStore() (*gorm.DB, error) {
	_ = godotenv.Load("../.env", "../.env.example")
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=prefer", username, password, host, port, dbname)
	return gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
}

func NewUserStore() (UserStore, error) {
	db, err := NewPostgresStore()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return &UserStoreImpl{db: db}, nil
}

func NewCardDeckStore() (CardDeckStore, error) {
	db, err := NewPostgresStore()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.CardSide{},
		&model.Card{},
		&model.Deck{},
		&model.UserActiveDecks{},
		&model.UserFavoriteDecks{})
	if err != nil {
		return nil, err
	}
	return &CardDeckStoreImpl{db: db}, nil
}

func NewCollaborationStore() (CollaborationStore, error) {
	db, err := NewPostgresStore()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Group{}, &model.GroupUserRole{})
	if err != nil {
		return nil, err
	}
	return &CollaborationStoreImpl{db: db}, nil
}

func NewSrsStore() (SrsStore, error) {
	db, err := NewPostgresStore()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Revlog{}, &model.UserCardBinding{})
	if err != nil {
		return nil, err
	}
	return &SrsStoreImpl{db: db}, nil
}

func (s *UserStoreImpl) RegisterNewUser(newUser *model.User) error {
	newUser.Email = strings.ToLower(newUser.Email)
	return s.db.Create(newUser).Error
}

func (s *UserStoreImpl) ModifyUser(user *model.User) error {
	user.Email = strings.ToLower(user.Email)
	return s.db.Save(&model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}).Error
}

func (s *UserStoreImpl) DeleteUser(user *model.User) error {
	return s.db.Delete(user).Error
}

func (s *UserStoreImpl) FindUserByEmail(email string) (user *model.User, err error) {
	email = strings.ToLower(email)
	if err = s.db.Where(&model.User{Email: email}).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoExistingUserWithEmail
	}
	return
}

func (s *UserStoreImpl) FindUserByID(userID string) (user *model.User, err error) {
	if err = s.db.Where(&model.User{ID: userID}).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoExistingUserWithEmail
	}
	return
}

func (s *CardDeckStoreImpl) PopulateDeckFavoriteAttribute(deck *model.Deck, userID string) error {
	var count int64
	if err := s.db.Model(&model.UserFavoriteDecks{}).Where(&model.UserFavoriteDecks{UserID: userID, DeckID: deck.ID}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		deck.IsFavorite = true
	}
	return nil
}

func (s *CardDeckStoreImpl) PopulateDeckActiveAttribute(deck *model.Deck, userID string) error {
	var count int64
	if err := s.db.Model(&model.UserActiveDecks{}).Where(&model.UserActiveDecks{UserID: userID, DeckID: deck.ID}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		deck.IsActive = true
	}
	return nil
}

func (s *CardDeckStoreImpl) FindDecks(userID string, groupID string, favoriteFilter bool, activeFilter bool) error {
	var decks []model.Deck
	if err := s.db.Where(&model.Deck{GroupID: groupID}).Find(&decks).Error; err != nil {
		return err
	}
	for _, deck := range decks {
		if err := s.PopulateDeckFavoriteAttribute(&deck, userID); err != nil {
			return err
		}
		if err := s.PopulateDeckActiveAttribute(&deck, userID); err != nil {
			return err
		}
	}
	return nil
}

func (s *CardDeckStoreImpl) FindDecksByGroupID(groupID string, userID string) (decks []model.Deck, err error) {
	if err = s.db.Where(&model.GroupUserRole{GroupID: groupID}).
		Find(&decks).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	for i := range decks {
		if err = s.PopulateDeckFavoriteAttribute(&decks[i], userID); err != nil {
			return
		}
		if err = s.PopulateDeckActiveAttribute(&decks[i], userID); err != nil {
			return
		}
	}
	return
}

func (s *CardDeckStoreImpl) FindDeckCards(deckID string) (cards []*model.Card, err error) {
	if err = s.db.Where(model.Card{DeckID: deckID}).
		Find(&cards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) FindPublicDecksByGroupID(groupID string) (decks []model.Deck, err error) {
	if err = s.db.Where(&model.GroupUserRole{GroupID: groupID}).
		Where(&model.Deck{DeckType: model.PublicDeckType}).
		Find(&decks).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) FindDeckByID(deckID string, userID string) (deck *model.Deck, err error) {
	if err = s.db.Where(&model.Deck{ID: deckID}).
		Preload("Cards").
		First(&deck).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	if err = s.PopulateDeckFavoriteAttribute(deck, userID); err != nil {
		return
	}
	if err = s.PopulateDeckActiveAttribute(deck, userID); err != nil {
		return
	}
	return
}

func (s *CardDeckStoreImpl) CreateDeck(newDeck *model.Deck) error {
	return s.db.Create(newDeck).Error
}

func (s *CardDeckStoreImpl) ModifyDeck(deck *model.Deck) (err error) {
	err = s.db.Save(&model.Deck{
		ID:        deck.ID,
		Name:      deck.Name,
		GroupID:   deck.GroupID,
		CreatedAt: deck.CreatedAt,
		DeckType:  deck.DeckType,
	}).Error
	return
}

func (s *CardDeckStoreImpl) DeleteDeck(deck *model.Deck) error {
	return s.db.Delete(deck).Error
}

func (s *CardDeckStoreImpl) FindCardByID(cardID string) (card *model.Card, err error) {
	if err = s.db.Where(&model.Card{ID: cardID}).
		Preload("Deck").
		First(&card).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) CreateCard(newCard *model.Card) error {
	return s.db.Create(newCard).Error
}

func (s *CardDeckStoreImpl) ModifyCard(card *model.Card) (err error) {
	err = s.db.Save(&model.Card{
		ID:              card.ID,
		DeckID:          card.DeckID,
		FirstCardSideID: card.FirstCardSideID,
		CreatedAt:       card.CreatedAt,
	}).Error
	return
}

func (s *CardDeckStoreImpl) DeleteCard(card *model.Card) error {
	return s.db.Delete(card).Error
}

func (s *CardDeckStoreImpl) FindCardSidesByCardID(cardID string) ([]model.CardSide, error) {
	card, err := s.FindCardByID(cardID)
	if err != nil {
		return nil, err
	}
	var cardSides []model.CardSide
	nextCardSideID := card.FirstCardSideID
	for finished := false; !finished; {
		cardSide, err := s.FindCardSideByID(nextCardSideID)
		if err != nil {
			return nil, err
		}
		cardSides = append(cardSides, *cardSide)
		if cardSide.NextCardSideID == "" {
			finished = true
		} else {
			nextCardSideID = cardSide.NextCardSideID
		}
	}
	return cardSides, nil
}

func (s *CardDeckStoreImpl) FindCardSideByID(cardSideID string) (cardSide *model.CardSide, err error) {
	if err = s.db.Where(&model.CardSide{ID: cardSideID}).
		Preload("Card").
		First(&cardSide).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) FindLastCardSideOfCardByID(cardID string) (cardSide *model.CardSide, err error) {
	if err = s.db.Where(&model.CardSide{CardID: cardID, NextCardSideID: ""}).
		Preload("Card").
		First(&cardSide).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) CreateCardSide(newCardSide *model.CardSide) error {
	return s.db.Create(newCardSide).Error
}

func (s *CardDeckStoreImpl) ModifyCardSide(cardSide *model.CardSide) error {
	return s.db.Save(&model.CardSide{
		ID:                 cardSide.ID,
		CardID:             cardSide.CardID,
		Header:             cardSide.Header,
		Description:        cardSide.Description,
		PreviousCardSideID: cardSide.PreviousCardSideID,
		NextCardSideID:     cardSide.NextCardSideID,
	}).Error
}

func (s *CardDeckStoreImpl) DeleteCardSide(cardSide *model.CardSide) error {
	return s.db.Delete(cardSide).Error
}

func (s *CardDeckStoreImpl) DeleteCardSidesOfCardByID(cardID string) error {
	return s.db.Where(&model.CardSide{CardID: cardID}).Delete(&model.CardSide{}).Error
}
func (s *CardDeckStoreImpl) FindFavoriteDecks(userID string) (decks []model.Deck, err error) {
	var userFavoriteDecks []model.UserFavoriteDecks
	err = s.db.
		Preload("Deck").
		Where(&model.UserFavoriteDecks{UserID: userID}).
		Find(&userFavoriteDecks).Error
	if err != nil {
		return
	}
	for _, deck := range userFavoriteDecks {
		deck.Deck.IsFavorite = true
		if err = s.PopulateDeckActiveAttribute(&deck.Deck, userID); err != nil {
			return
		}
		decks = append(decks, deck.Deck)
	}
	return
}

func (s *CardDeckStoreImpl) AddFavoriteDeck(userID string, deckID string) error {
	return s.db.Create(&model.UserFavoriteDecks{UserID: userID, DeckID: deckID}).Error
}

func (s *CardDeckStoreImpl) DeleteFavoriteDeck(userID string, deckID string) error {
	return s.db.Delete(&model.UserFavoriteDecks{UserID: userID, DeckID: deckID}).Error
}

func (s *CardDeckStoreImpl) FindActiveDecks(userID string) (decks []model.Deck, err error) {
	var userActiveDecks []model.UserActiveDecks
	err = s.db.
		Preload("Deck").
		Where(&model.UserActiveDecks{UserID: userID}).
		Find(&userActiveDecks).Error
	if err != nil {
		return
	}
	for _, deck := range userActiveDecks {
		deck.Deck.IsActive = true
		if err = s.PopulateDeckFavoriteAttribute(&deck.Deck, userID); err != nil {
			return
		}
		decks = append(decks, deck.Deck)
	}
	return
}

func (s *CardDeckStoreImpl) AddActiveDeck(userID string, deckID string) error {
	return s.db.Create(&model.UserActiveDecks{UserID: userID, DeckID: deckID, Algorithm: model.AlgoDynamicSRS}).Error
}

func (s *CardDeckStoreImpl) DeleteActiveDeck(userID string, deckID string) error {
	return s.db.Delete(&model.UserActiveDecks{UserID: userID, DeckID: deckID}).Error
}

func (s *CollaborationStoreImpl) FindGroupsByUserID(userID string) (groups []model.Group, err error) {
	if err = s.db.Joins("Join group_user_roles on group_user_roles.group_id = groups.id").
		Where("group_user_roles.user_id = ?", userID).
		Not("group_user_roles.role_type = ?", "requested").
		Not("group_user_roles.role_type = ?", "invited").
		Find(&groups).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupByID(groupID string) (group *model.Group, err error) {
	if err = s.db.Where(&model.Group{ID: groupID}).First(&group).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) CreateNewGroupWithAdmin(adminUserID string, newGroup *model.Group) error {
	if err := s.db.Create(newGroup).Error; err != nil {
		return err
	}
	return s.db.Create(&model.GroupUserRole{GroupID: newGroup.ID, UserID: adminUserID, RoleType: model.RoleAdmin}).Error
}

func (s *CollaborationStoreImpl) AddNewMemberToGroup(userID string, groupID string) error {
	return s.db.Create(&model.GroupUserRole{GroupID: groupID, UserID: userID, RoleType: model.RoleRead}).Error
}

func (s *CollaborationStoreImpl) AddInvitedUserToGroup(userID string, groupID string) error {
	return s.db.Create(&model.GroupUserRole{GroupID: groupID, UserID: userID, RoleType: model.RoleInvited}).Error
}

func (s *CollaborationStoreImpl) AddRequestingUserToGroup(userID string, groupID string) error {
	return s.db.Create(&model.GroupUserRole{GroupID: groupID, UserID: userID, RoleType: model.RoleRequested}).Error
}

func (s *CollaborationStoreImpl) PromoteUserToFullGroupMember(userID string, groupID string) error {
	return s.ModifyUserRole(userID, groupID, model.RoleRead)
}

func (s *CollaborationStoreImpl) ModifyUserRole(userID string, groupID string, role model.RoleType) error {
	var groupUserRole model.GroupUserRole
	if err := s.db.Where(&model.GroupUserRole{UserID: userID, GroupID: groupID}).
		First(&groupUserRole).Error; err != nil {
		return err
	}
	groupUserRole.RoleType = role
	if err := s.db.Save(&groupUserRole).Error; err != nil {
		return helper.ErrStoreInvalidGroupRoleForChange
	}
	return nil
}

func (s *CollaborationStoreImpl) RemoveUserFromGroup(userID string, groupID string) error {
	return s.db.Where(&model.GroupUserRole{UserID: userID, GroupID: groupID}).Delete(&model.GroupUserRole{}).Error
}

func (s *CollaborationStoreImpl) ModifyGroup(group *model.Group) (err error) {
	err = s.db.Save(&model.Group{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		IsDefault:   group.IsDefault,
		GroupType:   group.GroupType,
	}).Error
	return
}

func (s *CollaborationStoreImpl) DeleteGroup(group *model.Group) error {
	return s.db.Delete(group).Error
}

func (s *CollaborationStoreImpl) FindGroupUserRole(userID string, groupID string) (groupRole model.RoleType, err error) {
	var groupUser model.GroupUserRole
	if err = s.db.Where(&model.GroupUserRole{GroupID: groupID, UserID: userID}).
		First(&groupUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	groupRole = groupUser.RoleType
	return
}

func (s *CollaborationStoreImpl) FindGroupMemberRoles(groupID string) (groupMembers []model.GroupUserRole, err error) {
	if err = s.db.Where(&model.GroupUserRole{GroupID: groupID}).
		Not(&model.GroupUserRole{RoleType: model.RoleInvited}).
		Not(&model.GroupUserRole{RoleType: model.RoleRequested}).
		Find(&groupMembers).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupAdmins(groupID string) (groupMembers []model.GroupUserRole, err error) {
	if err = s.db.Where(&model.GroupUserRole{GroupID: groupID, RoleType: model.RoleAdmin}).
		Find(&groupMembers).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupRequestsByGroupID(groupID string) (requests []model.GroupUserRole, err error) {
	if err = s.db.Where(&model.GroupUserRole{GroupID: groupID, RoleType: model.RoleRequested}).
		Find(&requests).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupInvitationsByUserID(userID string) (invites []model.GroupUserRole, err error) {
	if err = s.db.Where(&model.GroupUserRole{UserID: userID, RoleType: model.RoleInvited}).
		Preload("Group").
		Find(&invites).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupInvitationsByGroupID(groupID string) (groupInvites []model.GroupUserRole, err error) {
	if err = s.db.Where(&model.GroupUserRole{GroupID: groupID, RoleType: model.RoleInvited}).
		Preload("Group").
		Find(&groupInvites).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) CreateRevlog(newRev *model.Revlog) error {
	return s.db.Create(&newRev).Error
}

func (s *SrsStoreImpl) FindCardBinding(userID string, cardID string) (userCardBinding *model.UserCardBinding, err error) {
	if err = s.db.Where(model.UserCardBinding{UserID: userID, CardID: cardID}).
		First(&userCardBinding).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) FindUserDeckCards(userID string, deckID string) (userCards []*model.UserCardBinding, err error) {
	if err = s.db.Where(model.UserCardBinding{UserID: userID, DeckID: deckID}).
		Find(&userCards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) FindUserCards(userID string) (userCards []*model.UserCardBinding, err error) {
	if err = s.db.Where(model.UserCardBinding{UserID: userID}).
		Find(&userCards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) CreateUserCard(newCard *model.UserCardBinding) error {
	return s.db.Create(newCard).Error
}

func (s *SrsStoreImpl) ModifyUserCard(userCard *model.UserCardBinding) (err error) {
	return s.db.Save(&model.UserCardBinding{
		ID:           userCard.ID,
		UserID:       userCard.UserID,
		CardID:       userCard.CardID,
		DeckID:       userCard.DeckID,
		Type:         userCard.Type,
		Due:          userCard.Due,
		LastInterval: userCard.LastInterval,
		Factor:       userCard.Factor,
	}).Error
}
