package store

import (
	"errors"
	"fmt"
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strings"

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
	err = db.AutoMigrate(&model.Group{}, &model.GroupUserRole{}, &model.GroupAdmission{})
	if err != nil {
		return nil, err
	}
	return &CollaborationStoreImpl{db: db}, nil
}

func (s *UserStoreImpl) RegisterNewUser(newUser *model.User) error {
	newUser.Email = strings.ToLower(newUser.Email)
	return s.db.Create(&newUser).Error
}

func (s *UserStoreImpl) FindUserByEmail(email string) (user *model.User, err error) {
	email = strings.ToLower(email)
	if err = s.db.Where(model.User{Email: email}).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoExistingUserWithEmail
	}
	return
}

func (s *UserStoreImpl) FindUserByID(userID string) (user *model.User, err error) {
	if err = s.db.Where(model.User{ID: userID}).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoExistingUserWithEmail
	}
	return
}

func (s *CardDeckStoreImpl) FindDecksByGroupID(groupID string) (decks []model.Deck, err error) {
	if err = s.db.Where(model.GroupUserRole{GroupID: groupID}).Find(&decks).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) FindDeckByID(deckID string) (deck *model.Deck, err error) {
	if err = s.db.Where(model.Deck{ID: deckID}).Preload("Cards").First(&deck).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) CreateDeck(newDeck *model.Deck) error {
	return s.db.Create(&newDeck).Error
}

func (s *CardDeckStoreImpl) ModifyDeck(deck *model.Deck) (err error) {
	err = s.db.Save(&model.Deck{
		ID:        deck.ID,
		Name:      deck.Name,
		GroupID:   deck.GroupID,
		CreatedAt: deck.CreatedAt,
	}).Error
	return
}

func (s *CardDeckStoreImpl) DeleteDeck(deck *model.Deck) error {
	return s.db.Delete(deck).Error
}

func (s *CardDeckStoreImpl) FindCardByID(cardID string) (card *model.Card, err error) {
	if err = s.db.Where(model.Card{ID: cardID}).Preload("Deck").Find(&card).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CardDeckStoreImpl) CreateCard(newCard *model.Card) error {
	return s.db.Create(&newCard).Error
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

func (s *CardDeckStoreImpl) DeleteCard(card *model.Card) error {
	return s.db.Delete(card).Error
}

func (s *CollaborationStoreImpl) FindGroupsByUserID(userID string) (groups []model.Group, err error) {
	if err = s.db.Joins("Join group_user_roles on group_user_roles.group_id = groups.id").
		Where("group_user_roles.user_id = ?", userID).
		Find(&groups).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupByID(groupID string) (group *model.Group, err error) {
	if err = s.db.Where(model.Group{ID: groupID}).First(&group).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) CreateNewGroupWithAdmin(adminUserID string, newGroup *model.Group) error {
	if err := s.db.Create(&newGroup).Error; err != nil {
		return err
	}
	return s.db.Create(model.GroupUserRole{GroupID: newGroup.ID, UserID: adminUserID, RoleType: model.RoleAdmin}).Error
}

func (s *CollaborationStoreImpl) AddNewMemberToGroup(userID string, groupID string) error {
	return s.db.Create(model.GroupUserRole{GroupID: groupID, UserID: userID, RoleType: model.RoleRead}).Error
}

func (s *CollaborationStoreImpl) ModifyGroup(group *model.Group) (err error) {
	err = s.db.Save(&model.Group{
		ID:        group.ID,
		Name:      group.Name,
		IsDefault: group.IsDefault,
		GroupType: group.GroupType,
	}).Error
	return
}

func (s *CollaborationStoreImpl) DeleteGroup(group *model.Group) error {
	return s.db.Delete(group).Error
}

func (s *CollaborationStoreImpl) GetGroupUserRole(userID string, groupID string) (groupRole model.RoleType, err error) {
	var groupUser model.GroupUserRole
	if err = s.db.Where(model.GroupUserRole{GroupID: groupID, UserID: userID}).First(&groupUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	groupRole = groupUser.RoleType
	return
}

func (s *CollaborationStoreImpl) GetGroupMemberRoles(groupID string) (groupMembers []model.GroupUserRole, err error) {
	if err = s.db.Where(model.GroupUserRole{GroupID: groupID}).Find(&groupMembers).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) CreateNewGroupAdmission(newAdmission *model.GroupAdmission) error {
	return s.db.Create(&newAdmission).Error
}

func (s *CollaborationStoreImpl) FindGroupRequestsByGroupID(groupID string) (groupAdmissions []model.GroupAdmission, err error) {
	if err = s.db.Where(model.GroupAdmission{GroupID: groupID, AdmissionStatus: model.Requested}).Find(&groupAdmissions).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupInvitationsByUserID(userID string) (groupAdmissions []model.GroupAdmission, err error) {
	if err = s.db.Where(model.GroupAdmission{UserID: userID, AdmissionStatus: model.Invited}).Preload("Group").Find(&groupAdmissions).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupAdmissionByUserAndGroupID(userID string, groupID string) (groupAdmission *model.GroupAdmission, err error) {
	if err = s.db.Where(model.GroupAdmission{UserID: userID, GroupID: groupID}).First(&groupAdmission).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) FindGroupAdmissionByID(admissionID string) (groupAdmission *model.GroupAdmission, err error) {
	if err = s.db.Where(model.GroupAdmission{ID: admissionID}).First(&groupAdmission).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = helper.ErrStoreNoEntryWithID
	}
	return
}

func (s *CollaborationStoreImpl) DeleteGroupAdmission(admission *model.GroupAdmission) error {
	return s.db.Delete(admission).Error
}
