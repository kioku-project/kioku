package store

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kioku-project/kioku/pkg/model"
	"github.com/joho/godotenv"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
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

	err = db.AutoMigrate(&model.Card{}, &model.Deck{}, &model.Group{}, &model.User{}, &model.GroupUser{})
	if err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) FindUserByEmail(email string) (*model.User, error) {
	var usr model.User
	err := s.db.Where("email = ?", email).First(&usr).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (s *PostgresStore) RegisterNewUser(newUser *model.User) (error) {
	s.db.Create(&newUser)
	return nil
}
