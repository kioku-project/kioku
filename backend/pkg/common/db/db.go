package db

import (
	"fmt"
	"github.com/apex/log"
	"github.com/kioku-project/kioku/pkg/card"
	"github.com/kioku-project/kioku/pkg/deck"
	"github.com/kioku-project/kioku/pkg/group"
	"github.com/kioku-project/kioku/pkg/groupUser"
	"github.com/kioku-project/kioku/pkg/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitializeDB() *gorm.DB {
	var (
		host     = "db"
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		username, password, host, dbname)
	var err error
	DB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("Could not open database connection")
	}

	// initialize and migrate tables
	if err = DB.AutoMigrate(&card.Card{}, &deck.Deck{}, &group.Group{}, &user.User{}, &groupUser.GroupUser{}); err != nil {
		log.WithError(err).Fatal("Error while creating/migrating database tables:")
	}

	return DB
}
