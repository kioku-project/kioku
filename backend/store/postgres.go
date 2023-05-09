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

func NewPostgresStore() (*gorm.DB, error) {
	_ = godotenv.Load("../.env", "../.env.example")
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewUserStore() (*PostgresStore, error) {
	db, err := NewPostgresStore()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func NewCardDeckStore() (*PostgresStore, error) {
	db, err := NewPostgresStore()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Card{}, &model.Deck{})
	if err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func NewCollaborationStore() (*PostgresStore, error) {
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
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) FindUserByEmail(email string) (*model.User, error) {
	var usr model.User
	err := s.db.Where(model.User{Email: email}).First(&usr).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (s *PostgresStore) RegisterNewUser(newUser *model.User) error {
	rsp := s.db.Create(&newUser)
	return rsp.Error
}

func (s *PostgresStore) CreateDeck(newDeck *model.Deck) error {
	rsp := s.db.Create(&newDeck)
	return rsp.Error
}

func (s *PostgresStore) FindDeckByPublicID(publicID string) (*model.Deck, error) {
	var deck model.Deck
	err := s.db.Where(model.Deck{PublicID: publicID}).First(&deck).Error
	if err != nil {
		return nil, err
	}
	return &deck, nil
}

func (s *PostgresStore) CreateCard(newCard *model.Card) error {
	rsp := s.db.Create(&newCard)
	return rsp.Error
}

func (s *PostgresStore) CreateNewGroupWithAdmin(adminUserID uint, newGroup *model.Group) error {
	rsp := s.db.Create(&newGroup)
	if rsp.Error != nil {
		return rsp.Error
	}
	rsp = s.db.Create(model.GroupUserRole{GroupID: newGroup.ID, UserID: adminUserID, RoleType: model.RoleAdmin})
	return rsp.Error
}

func (s *PostgresStore) FindGroupByPublicID(publicID string) (*model.Group, error) {
	var group model.Group
	err := s.db.Where(model.Group{PublicID: publicID}).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *PostgresStore) GetGroupUserRole(userID uint, groupID uint) (*model.RoleType, error) {
	var groupUser model.GroupUserRole
	err := s.db.Where(model.GroupUserRole{GroupID: groupID, UserID: userID}).First(&groupUser).Error
	if err != nil {
		return nil, err
	}
	return &groupUser.RoleType, nil
}
