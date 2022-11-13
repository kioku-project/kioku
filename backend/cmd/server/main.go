package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/card"
	"github.com/kioku-project/kioku/pkg/deck"
	"github.com/kioku-project/kioku/pkg/group"
	"github.com/kioku-project/kioku/pkg/groupUser"
	"github.com/kioku-project/kioku/pkg/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func main() {

	_ = godotenv.Load()

	initializeDB()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3001")
}

func initializeDB() {

	var (
		_      = os.Getenv("DB_HOST")
		_      = os.Getenv("DB_PORT")
		_      = os.Getenv("DB_USER")
		_      = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
	)

	connStr := fmt.Sprintf("%s.sqlite", dbname)
	log.Info(dbname)
	var err error
	DB, err := gorm.Open(sqlite.Open(connStr), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("Could not open database connection")
	}

	// initialize and migrate tables
	if err = DB.AutoMigrate(&card.Card{}, &deck.Deck{}, &group.Group{}, &user.User{}, &groupUser.GroupUser{}); err != nil {
		log.WithError(err).Fatal("Error while creating/migrating database tables:")
	}

}
