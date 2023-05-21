package store

import (
	"errors"
	"fmt"
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/model"
)

type GormStore struct {
	db *gorm.DB
}

type UserStoreImpl GormStore

type CardDeckStoreImpl GormStore

type CollaborationStoreImpl GormStore

func NewPostgresStore() (*gorm.DB, error) {
	_ = godotenv.Load("../.env", "../.env.example")
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)
	return gorm.Open(postgres.Open(connStr), &gorm.Config{})
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
	err = db.AutoMigrate(&model.Card{}, &model.Deck{})
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
	err = db.SetupJoinTable(&model.User{}, "Groups", &model.GroupUserRole{})
	if err != nil {
		return nil, err
	}
	err = db.SetupJoinTable(&model.Group{}, "Users", &model.GroupUserRole{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Group{})
	if err != nil {
		return nil, err
	}
	return &CollaborationStoreImpl{db: db}, nil
}

func (s *UserStoreImpl) FindUserByEmail(email string) (user *model.User, err error) {
	err = s.db.Where(model.User{Email: email}).Preload("Groups").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helper.ErrStoreNoExistingUserWithEmail
		}
	}
	return
}

func (s *UserStoreImpl) RegisterNewUser(newUser *model.User) (err error) {
	err = s.db.Create(&newUser).Error
	return
}

func (s *CardDeckStoreImpl) CreateDeck(newDeck *model.Deck) (err error) {
	err = s.db.Create(&newDeck).Error
	return
}

func (s *CardDeckStoreImpl) FindDeckByID(deckID string) (deck *model.Deck, err error) {
	err = s.db.Where(model.Deck{ID: deckID}).Preload("Group").Preload("Cards").First(&deck).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helper.ErrStoreNoEntryWithID
		}
	}
	return
}

func (s *CardDeckStoreImpl) CreateCard(newCard *model.Card) (err error) {
	err = s.db.Create(&newCard).Error
	return
}

func (s *CardDeckStoreImpl) FindDecksByGroupID(groupID string) (decks []model.Deck, err error) {
	err = s.db.Where(model.GroupUserRole{GroupID: groupID}).Find(&decks).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helper.ErrStoreNoEntryWithID
		}
	}
	return
}

func (s *CardDeckStoreImpl) FindCardByID(cardID string) (card *model.Card, err error) {
	err = s.db.Where(model.Card{ID: cardID}).Preload("Deck").Find(&card).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helper.ErrStoreNoEntryWithID
		}
	}
	return
}

func (s *CardDeckStoreImpl) DeleteCard(card *model.Card) (err error) {
	err = s.db.Delete(card).Error
	return
}

func (s *CardDeckStoreImpl) DeleteDeck(deck *model.Deck) (err error) {
	err = s.db.Delete(deck).Error
	return
}

func (s *CardDeckStoreImpl) ModifyCard(card *model.Card) (err error) {
	err = s.db.Save(&model.Card{
		ID:        card.ID,
		DeckID:    card.DeckID,
		Frontside: card.Frontside,
		Backside:  card.Backside,
	}).Error
	return
}

func (s *CardDeckStoreImpl) ModifyDeck(deck *model.Deck) (err error) {
	err = s.db.Save(&model.Deck{
		ID:      deck.ID,
		Name:    deck.Name,
		GroupID: deck.GroupID,
	}).Error
	return
}

func (s *CollaborationStoreImpl) CreateNewGroupWithAdmin(adminUserID string, newGroup *model.Group) (err error) {
	err = s.db.Create(&newGroup).Error
	if err != nil {
		return
	}
	err = s.db.Create(model.GroupUserRole{GroupID: newGroup.ID, UserID: adminUserID, RoleType: model.RoleAdmin}).Error
	return
}

func (s *CollaborationStoreImpl) FindGroupByID(groupID string) (group *model.Group, err error) {
	err = s.db.Where(model.Group{ID: groupID}).First(&group).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helper.ErrStoreNoEntryWithID
		}
	}
	return
}

func (s *CollaborationStoreImpl) GetGroupUserRole(userID string, groupID string) (model.RoleType, error) {
	var groupUser model.GroupUserRole
	err := s.db.Where(model.GroupUserRole{GroupID: groupID, UserID: userID}).First(&groupUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", helper.ErrStoreNoEntryWithID
		}
		return "", err
	}
	return groupUser.RoleType, nil
}

func (s *CollaborationStoreImpl) FindGroupsByUserID(userID string) (groups []model.Group, err error) {
	err = s.db.Joins("Join group_user_roles on group_user_roles.group_id = groups.id").
		Where("group_user_roles.user_id = ?", userID).
		Find(&groups).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helper.ErrStoreNoEntryWithID
		}
	}
	return
}

func (s *CollaborationStoreImpl) DeleteGroup(group *model.Group) (err error) {
	err = s.db.Delete(group).Error
	return
}

func (s *CollaborationStoreImpl) ModifyGroup(group *model.Group) (err error) {
	err = s.db.Save(&model.Group{
		ID:        group.ID,
		Name:      group.Name,
		IsDefault: group.IsDefault,
	}).Error
	return
}
