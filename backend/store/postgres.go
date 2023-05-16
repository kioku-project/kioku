package store

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/model"
)

type PostgresStore struct {
	db *gorm.DB
}

type UserStoreImpl struct {
	PostgresStore
}

type CardDeckStoreImpl struct {
	PostgresStore
}

type CollaborationStoreImpl struct {
	PostgresStore
}

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
	return &UserStoreImpl{PostgresStore: PostgresStore{db: db}}, nil
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
	return &CardDeckStoreImpl{PostgresStore: PostgresStore{db: db}}, nil
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
	return &CollaborationStoreImpl{PostgresStore: PostgresStore{db: db}}, nil
}

func (s *UserStoreImpl) FindUserByEmail(email string) (user *model.User, err error) {
	err = s.db.Where(model.User{Email: email}).First(&user).Error
	return
}

func (s *UserStoreImpl) RegisterNewUser(newUser *model.User) error {
	rsp := s.db.Create(&newUser)
	return rsp.Error
}

func (s *CardDeckStoreImpl) CreateDeck(newDeck *model.Deck) error {
	rsp := s.db.Create(&newDeck)
	return rsp.Error
}

func (s *CardDeckStoreImpl) FindDeckByPublicID(publicID string) (deck *model.Deck, err error) {
	err = s.db.Where(model.Deck{PublicID: publicID}).Preload("Cards").First(&deck).Error
	return
}

func (s *CardDeckStoreImpl) CreateCard(newCard *model.Card) error {
	rsp := s.db.Create(&newCard)
	return rsp.Error
}

func (s *CardDeckStoreImpl) FindDecksByGroupID(groupID uint) (decks []model.Deck, err error) {
	err = s.db.Where(model.GroupUserRole{GroupID: groupID}).Find(&decks).Error
	return
}

func (s *CollaborationStoreImpl) CreateNewGroupWithAdmin(adminUserID uint, newGroup *model.Group) error {
	rsp := s.db.Create(&newGroup)
	if rsp.Error != nil {
		return rsp.Error
	}
	rsp = s.db.Create(model.GroupUserRole{GroupID: newGroup.ID, UserID: adminUserID, RoleType: model.RoleAdmin})
	return rsp.Error
}

func (s *PostgresStore) FindGroupByPublicID(publicID string) (group *model.Group, err error) {
	err = s.db.Where(model.Group{PublicID: publicID}).First(&group).Error
	return
}

func (s *CollaborationStoreImpl) GetGroupUserRole(userID uint, groupID uint) (*model.RoleType, error) {
	var groupUser model.GroupUserRole
	err := s.db.Where(model.GroupUserRole{GroupID: groupID, UserID: userID}).First(&groupUser).Error
	if err != nil {
		return nil, err
	}
	return &groupUser.RoleType, nil
}

func (s *CollaborationStoreImpl) FindGroupsByUserID(userID uint) (groups []model.Group, err error) {
	err = s.db.Joins("Join group_user_roles on group_user_roles.group_id = groups.id").
		Where("group_user_roles.user_id = ?", userID).
		Find(&groups).Error
	return
}
