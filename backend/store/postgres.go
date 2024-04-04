package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"go.opentelemetry.io/otel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/model"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

type GormStore struct {
	db *gorm.DB
}

type UserStoreImpl GormStore

type CardDeckStoreImpl GormStore

type CollaborationStoreImpl GormStore

type SrsStoreImpl GormStore

type NotificationStoreImpl GormStore

func NewPostgresStore(ctx context.Context) (*gorm.DB, error) {
	_ = godotenv.Load("../.env", "../.env.example")
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	logger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      false,
		},
	)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=prefer", username, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		TranslateError: true,
		Logger:         logger,
	})
	if err != nil {
		return nil, err
	}

	if err = db.Use(tracing.NewPlugin(tracing.WithDBName("kioku"))); err != nil {
		return nil, err
	}

	tracer := otel.Tracer("gorm.io/plugin/opentelemetry")

	ctx, span := tracer.Start(ctx, "root")
	defer span.End()

	return db.WithContext(ctx), nil
}

func NewUserStore(ctx context.Context) (UserStore, error) {
	db, err := NewPostgresStore(ctx)
	if err != nil {
		return nil, err
	}
	err = db.WithContext(ctx).AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return &UserStoreImpl{db: db}, nil
}

func NewCardDeckStore(ctx context.Context) (CardDeckStore, error) {
	db, err := NewPostgresStore(ctx)
	if err != nil {
		return nil, err
	}
	err = db.WithContext(ctx).AutoMigrate(&model.CardSide{},
		&model.Card{},
		&model.Deck{},
		&model.UserActiveDecks{},
		&model.UserFavoriteDecks{})
	if err != nil {
		return nil, err
	}
	return &CardDeckStoreImpl{db: db}, nil
}

func NewCollaborationStore(ctx context.Context) (CollaborationStore, error) {
	db, err := NewPostgresStore(ctx)
	if err != nil {
		return nil, err
	}
	err = db.WithContext(ctx).AutoMigrate(&model.Group{}, &model.GroupUserRole{})
	if err != nil {
		return nil, err
	}
	return &CollaborationStoreImpl{db: db}, nil
}

func NewSrsStore(ctx context.Context) (SrsStore, error) {
	db, err := NewPostgresStore(ctx)
	if err != nil {
		return nil, err
	}
	err = db.WithContext(ctx).AutoMigrate(&model.Revlog{}, &model.UserCardBinding{})
	if err != nil {
		return nil, err
	}
	return &SrsStoreImpl{db: db}, nil
}

func (s *UserStoreImpl) RegisterNewUser(ctx context.Context, newUser *model.User) error {
	newUser.Email = strings.ToLower(newUser.Email)
	return s.db.WithContext(ctx).Create(newUser).Error
}

func (s *UserStoreImpl) ModifyUser(ctx context.Context, user *model.User) error {
	user.Email = strings.ToLower(user.Email)
	return s.db.WithContext(ctx).Save(&model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}).Error
}

func (s *UserStoreImpl) DeleteUser(ctx context.Context, user *model.User) error {
	return s.db.WithContext(ctx).Delete(user).Error
}

func (s *UserStoreImpl) FindUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	email = strings.ToLower(email)
	if err = s.db.WithContext(ctx).Where(&model.User{Email: email}).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoExistingUserWithEmail
	}
	return
}

func (s *UserStoreImpl) FindUserByID(ctx context.Context, userID string) (user *model.User, err error) {
	if err = s.db.WithContext(ctx).Where(&model.User{ID: userID}).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoExistingUserWithEmail
	}
	return
}

func (s *CardDeckStoreImpl) PopulateDeckFavoriteAttribute(ctx context.Context, deck *model.Deck, userID string) error {
	var count int64
	if err := s.db.WithContext(ctx).Model(&model.UserFavoriteDecks{}).Where(&model.UserFavoriteDecks{UserID: userID, DeckID: deck.ID}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		deck.IsFavorite = true
	}
	return nil
}

func (s *CardDeckStoreImpl) PopulateDeckActiveAttribute(ctx context.Context, deck *model.Deck, userID string) error {
	var userActiveDeck model.UserActiveDecks
	if err := s.db.WithContext(ctx).Model(&model.UserActiveDecks{}).Where(&model.UserActiveDecks{UserID: userID, DeckID: deck.ID}).Find(&userActiveDeck).Error; err != nil {
		return err
	}
	if userActiveDeck.UserID == userID {
		deck.IsActive = true
		deck.Algorithm = userActiveDeck.Algorithm
	}
	return nil
}

func (s *CardDeckStoreImpl) FindDecksByGroupID(ctx context.Context, groupID string, userID string) (decks []model.Deck, err error) {
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{GroupID: groupID}).
		Find(&decks).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	for i := range decks {
		if err = s.PopulateDeckFavoriteAttribute(ctx, &decks[i], userID); err != nil {
			return
		}
		if err = s.PopulateDeckActiveAttribute(ctx, &decks[i], userID); err != nil {
			return
		}
	}
	return
}

func (s *CardDeckStoreImpl) FindDeckCards(ctx context.Context, deckID string) (cards []*model.Card, err error) {
	if err = s.db.WithContext(ctx).Where(model.Card{DeckID: deckID}).
		Find(&cards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) FindPublicDecksByGroupID(ctx context.Context, groupID string) (decks []model.Deck, err error) {
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{GroupID: groupID}).
		Where(&model.Deck{DeckType: model.PublicDeckType}).
		Find(&decks).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) FindDeckByID(ctx context.Context, deckID string, userID string) (deck *model.Deck, err error) {
	if err = s.db.WithContext(ctx).Where(&model.Deck{ID: deckID}).
		Preload("Cards").
		First(&deck).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	if err = s.PopulateDeckFavoriteAttribute(ctx, deck, userID); err != nil {
		return
	}
	if err = s.PopulateDeckActiveAttribute(ctx, deck, userID); err != nil {
		return
	}
	return
}

func (s *CardDeckStoreImpl) CreateDeck(ctx context.Context, newDeck *model.Deck) error {
	return s.db.WithContext(ctx).Create(newDeck).Error
}

func (s *CardDeckStoreImpl) ModifyDeck(ctx context.Context, deck *model.Deck) (err error) {
	err = s.db.WithContext(ctx).Save(&model.Deck{
		ID:          deck.ID,
		Name:        deck.Name,
		GroupID:     deck.GroupID,
		CreatedAt:   deck.CreatedAt,
		DeckType:    deck.DeckType,
		Description: deck.Description,
	}).Error
	return
}

func (s *CardDeckStoreImpl) DeleteDeck(ctx context.Context, deck *model.Deck) error {
	return s.db.WithContext(ctx).Delete(deck).Error
}

func (s *CardDeckStoreImpl) FindCardByID(ctx context.Context, cardID string) (card *model.Card, err error) {
	if err = s.db.WithContext(ctx).Where(&model.Card{ID: cardID}).
		Preload("Deck").
		First(&card).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) CreateCard(ctx context.Context, newCard *model.Card) error {
	return s.db.WithContext(ctx).Create(newCard).Error
}

func (s *CardDeckStoreImpl) ModifyCard(ctx context.Context, card *model.Card) (err error) {
	err = s.db.WithContext(ctx).Save(&model.Card{
		ID:              card.ID,
		DeckID:          card.DeckID,
		FirstCardSideID: card.FirstCardSideID,
		CreatedAt:       card.CreatedAt,
	}).Error
	return
}

func (s *CardDeckStoreImpl) DeleteCard(ctx context.Context, card *model.Card) error {
	return s.db.WithContext(ctx).Delete(card).Error
}

func (s *CardDeckStoreImpl) FindCardSidesByCardID(ctx context.Context, cardID string) ([]model.CardSide, error) {
	card, err := s.FindCardByID(ctx, cardID)
	if err != nil {
		return nil, err
	}
	var cardSides []model.CardSide
	nextCardSideID := card.FirstCardSideID
	for finished := false; !finished; {
		cardSide, err := s.FindCardSideByID(ctx, nextCardSideID)
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

func (s *CardDeckStoreImpl) FindCardSideByID(ctx context.Context, cardSideID string) (cardSide *model.CardSide, err error) {
	if err = s.db.WithContext(ctx).Where(&model.CardSide{ID: cardSideID}).
		Preload("Card").
		First(&cardSide).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) FindLastCardSideOfCardByID(ctx context.Context, cardID string) (cardSide *model.CardSide, err error) {
	if err = s.db.WithContext(ctx).Where(&model.CardSide{CardID: cardID, NextCardSideID: ""}).
		Preload("Card").
		First(&cardSide).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) CreateCardSide(ctx context.Context, newCardSide *model.CardSide) error {
	return s.db.WithContext(ctx).Create(newCardSide).Error
}

func (s *CardDeckStoreImpl) ModifyCardSide(ctx context.Context, cardSide *model.CardSide) error {
	return s.db.WithContext(ctx).Save(&model.CardSide{
		ID:                 cardSide.ID,
		CardID:             cardSide.CardID,
		Header:             cardSide.Header,
		Description:        cardSide.Description,
		PreviousCardSideID: cardSide.PreviousCardSideID,
		NextCardSideID:     cardSide.NextCardSideID,
	}).Error
}

func (s *CardDeckStoreImpl) DeleteCardSide(ctx context.Context, cardSide *model.CardSide) error {
	return s.db.WithContext(ctx).Delete(cardSide).Error
}

func (s *CardDeckStoreImpl) DeleteCardSidesOfCardByID(ctx context.Context, cardID string) error {
	return s.db.WithContext(ctx).Where(&model.CardSide{CardID: cardID}).Delete(&model.CardSide{}).Error
}
func (s *CardDeckStoreImpl) FindFavoriteDecks(ctx context.Context, userID string) (decks []model.Deck, err error) {
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
		if err = s.PopulateDeckActiveAttribute(ctx, &deck.Deck, userID); err != nil {
			return
		}
		decks = append(decks, deck.Deck)
	}
	return
}

func (s *CardDeckStoreImpl) AddFavoriteDeck(ctx context.Context, userID string, deckID string) error {
	return s.db.WithContext(ctx).Create(&model.UserFavoriteDecks{UserID: userID, DeckID: deckID}).Error
}

func (s *CardDeckStoreImpl) DeleteFavoriteDeck(ctx context.Context, userID string, deckID string) error {
	return s.db.WithContext(ctx).Delete(&model.UserFavoriteDecks{UserID: userID, DeckID: deckID}).Error
}

func (s *CardDeckStoreImpl) FindActiveDecks(ctx context.Context, userID string) (decks []model.Deck, err error) {
	var userActiveDecks []model.UserActiveDecks
	err = s.db.
		WithContext(ctx).
		Preload("Deck").
		Where(&model.UserActiveDecks{UserID: userID}).
		Find(&userActiveDecks).Error
	if err != nil {
		return
	}
	for _, deck := range userActiveDecks {
		deck.Deck.IsActive = true
		deck.Deck.Algorithm = deck.Algorithm
		if err = s.PopulateDeckFavoriteAttribute(ctx, &deck.Deck, userID); err != nil {
			return
		}
		decks = append(decks, deck.Deck)
	}
	return
}

func (s *CardDeckStoreImpl) AddActiveDeck(ctx context.Context, userID string, deckID string) error {
	return s.db.WithContext(ctx).Create(&model.UserActiveDecks{UserID: userID, DeckID: deckID, Algorithm: model.AlgoDynamicSRS, NewCardsPerDay: 5}).Error
}

func (s *CardDeckStoreImpl) DeleteActiveDeck(ctx context.Context, userID string, deckID string) error {
	return s.db.WithContext(ctx).Delete(&model.UserActiveDecks{UserID: userID, DeckID: deckID}).Error
}

func (s *CollaborationStoreImpl) FindGroupsByUserID(ctx context.Context, userID string) (groups []model.Group, err error) {
	if err = s.db.WithContext(ctx).Joins("Join group_user_roles on group_user_roles.group_id = groups.id").
		Where("group_user_roles.user_id = ?", userID).
		Not("group_user_roles.role_type = ?", "requested").
		Not("group_user_roles.role_type = ?", "invited").
		Find(&groups).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupByID(ctx context.Context, groupID string) (group *model.Group, err error) {
	if err = s.db.WithContext(ctx).Where(&model.Group{ID: groupID}).First(&group).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) CreateNewGroupWithAdmin(ctx context.Context, adminUserID string, newGroup *model.Group) error {
	if err := s.db.WithContext(ctx).Create(newGroup).Error; err != nil {
		return err
	}
	return s.db.WithContext(ctx).Create(&model.GroupUserRole{GroupID: newGroup.ID, UserID: adminUserID, RoleType: model.RoleAdmin}).Error
}

func (s *CollaborationStoreImpl) AddNewMemberToGroup(ctx context.Context, userID string, groupID string) error {
	return s.db.WithContext(ctx).Create(&model.GroupUserRole{GroupID: groupID, UserID: userID, RoleType: model.RoleRead}).Error
}

func (s *CollaborationStoreImpl) AddInvitedUserToGroup(ctx context.Context, userID string, groupID string) error {
	return s.db.WithContext(ctx).Create(&model.GroupUserRole{GroupID: groupID, UserID: userID, RoleType: model.RoleInvited}).Error
}

func (s *CollaborationStoreImpl) AddRequestingUserToGroup(ctx context.Context, userID string, groupID string) error {
	return s.db.WithContext(ctx).Create(&model.GroupUserRole{GroupID: groupID, UserID: userID, RoleType: model.RoleRequested}).Error
}

func (s *CollaborationStoreImpl) PromoteUserToFullGroupMember(ctx context.Context, userID string, groupID string) error {
	return s.ModifyUserRole(ctx, userID, groupID, model.RoleRead)
}

func (s *CollaborationStoreImpl) ModifyUserRole(ctx context.Context, userID string, groupID string, role model.RoleType) error {
	var groupUserRole model.GroupUserRole
	if err := s.db.WithContext(ctx).Where(&model.GroupUserRole{UserID: userID, GroupID: groupID}).
		First(&groupUserRole).Error; err != nil {
		return err
	}
	groupUserRole.RoleType = role
	if err := s.db.WithContext(ctx).Save(&groupUserRole).Error; err != nil {
		return helper.ErrStoreInvalidGroupRoleForChange
	}
	return nil
}

func (s *CollaborationStoreImpl) RemoveUserFromGroup(ctx context.Context, userID string, groupID string) error {
	return s.db.WithContext(ctx).Where(&model.GroupUserRole{UserID: userID, GroupID: groupID}).Delete(&model.GroupUserRole{}).Error
}

func (s *CollaborationStoreImpl) ModifyGroup(ctx context.Context, group *model.Group) (err error) {
	err = s.db.WithContext(ctx).Save(&model.Group{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		IsDefault:   group.IsDefault,
		GroupType:   group.GroupType,
	}).Error
	return
}

func (s *CollaborationStoreImpl) DeleteGroup(ctx context.Context, group *model.Group) error {
	return s.db.WithContext(ctx).Delete(group).Error
}

func (s *CollaborationStoreImpl) FindGroupUserRole(ctx context.Context, userID string, groupID string) (groupRole model.RoleType, err error) {
	var groupUser model.GroupUserRole
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{GroupID: groupID, UserID: userID}).
		First(&groupUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	groupRole = groupUser.RoleType
	return
}

func (s *CollaborationStoreImpl) FindGroupMemberRoles(ctx context.Context, groupID string) (groupMembers []model.GroupUserRole, err error) {
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{GroupID: groupID}).
		Not(&model.GroupUserRole{RoleType: model.RoleInvited}).
		Not(&model.GroupUserRole{RoleType: model.RoleRequested}).
		Find(&groupMembers).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupAdmins(ctx context.Context, groupID string) (groupMembers []model.GroupUserRole, err error) {
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{GroupID: groupID, RoleType: model.RoleAdmin}).
		Find(&groupMembers).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupRequestsByGroupID(ctx context.Context, groupID string) (requests []model.GroupUserRole, err error) {
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{GroupID: groupID, RoleType: model.RoleRequested}).
		Find(&requests).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupInvitationsByUserID(ctx context.Context, userID string) (invites []model.GroupUserRole, err error) {
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{UserID: userID, RoleType: model.RoleInvited}).
		Preload("Group").
		Find(&invites).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupInvitationsByGroupID(ctx context.Context, groupID string) (groupInvites []model.GroupUserRole, err error) {
	if err = s.db.WithContext(ctx).Where(&model.GroupUserRole{GroupID: groupID, RoleType: model.RoleInvited}).
		Preload("Group").
		Find(&groupInvites).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) CreateRevlog(ctx context.Context, newRev *model.Revlog) error {
	return s.db.WithContext(ctx).Create(&newRev).Error
}

func (s *SrsStoreImpl) FindCardBinding(ctx context.Context, userID string, cardID string) (userCardBinding *model.UserCardBinding, err error) {
	if err = s.db.WithContext(ctx).Where(model.UserCardBinding{UserID: userID, CardID: cardID}).
		First(&userCardBinding).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) FindUserDeckCards(ctx context.Context, userID string, deckID string) (userCards []*model.UserCardBinding, err error) {
	if err = s.db.WithContext(ctx).Where(model.UserCardBinding{UserID: userID, DeckID: deckID}).
		Find(&userCards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) FindUserDeckDueCards(ctx context.Context, userID string, deckID string) (userCards []*model.UserCardBinding, err error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrow := today.Add(24 * time.Hour)

	if err = s.db.WithContext(ctx).Where(model.UserCardBinding{UserID: userID, DeckID: deckID}).
		Not("user_card_bindings.due = ?", 0).
		Where("user_card_bindings.due <= ?", tomorrow.Unix()).
		Find(&userCards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) FindUserDeckNewCards(ctx context.Context, userID string, deckID string) (userCards []*model.UserCardBinding, err error) {
	if err = s.db.WithContext(ctx).
		Where(model.UserCardBinding{UserID: userID, DeckID: deckID}).
		Where("user_card_bindings.due = ?", 0).
		Find(&userCards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) FindUserDeckNewCardsLearnedToday(ctx context.Context, userID string, deckID string) (count int64, err error) {

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrow := today.Add(24 * time.Hour)

	err = s.db.Debug().WithContext(ctx).
		Model(&model.Revlog{}).
		Select("revlogs.id").
		Joins("LEFT JOIN user_card_bindings ON revlogs.card_id = user_card_bindings.card_id").
		Where(&model.Revlog{UserID: userID}).
		Where("user_card_bindings.deck_id = ?", deckID).
		Where("revlogs.due = ?", 0).
		Where("date >= ?", today.Unix()).
		Where("date <= ?", tomorrow.Unix()).
		Count(&count).Error

	return count, err
}

func (s *SrsStoreImpl) FindUserCards(ctx context.Context, userID string) (userCards []*model.UserCardBinding, err error) {
	if err = s.db.WithContext(ctx).Where(model.UserCardBinding{UserID: userID}).
		Find(&userCards).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *SrsStoreImpl) CreateUserCard(ctx context.Context, newCard *model.UserCardBinding) error {
	return s.db.WithContext(ctx).Create(newCard).Error
}

func (s *SrsStoreImpl) ModifyUserCard(ctx context.Context, userCard *model.UserCardBinding) (err error) {
	return s.db.WithContext(ctx).Save(&model.UserCardBinding{
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

func NewNotificationStore(ctx context.Context) (NotificationStore, error) {
	db, err := NewPostgresStore(ctx)
	if err != nil {
		return nil, err
	}
	err = db.WithContext(ctx).AutoMigrate(&model.PushSubscription{})
	if err != nil {
		return nil, err
	}
	return &NotificationStoreImpl{db: db}, nil
}

func (s *NotificationStoreImpl) CreatePushSubscription(ctx context.Context, newSubscription *model.PushSubscription) error {
	return s.db.WithContext(ctx).Create(newSubscription).Error
}

func (s *NotificationStoreImpl) FindAllPushSubscriptions(ctx context.Context) (subscriptions []*model.PushSubscription, err error) {
	return subscriptions, s.db.WithContext(ctx).Find(&subscriptions).Error
}

func (s *NotificationStoreImpl) FindPushSubscriptionByID(ctx context.Context, subscriptionID string) (subscription *model.PushSubscription, err error) {
	if err = s.db.WithContext(ctx).Where(&model.PushSubscription{ID: subscriptionID}).First(&subscription).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *NotificationStoreImpl) DeletePushSubscription(ctx context.Context, subscription *model.PushSubscription) error {
	return s.db.WithContext(ctx).Delete(subscription).Error
}
func (s *NotificationStoreImpl) FindPushSubscriptionsByUserID(ctx context.Context, userID string) (subscriptions []*model.PushSubscription, err error) {
	return subscriptions, s.db.WithContext(ctx).Where(&model.PushSubscription{UserID: userID}).Find(&subscriptions).Error
}
