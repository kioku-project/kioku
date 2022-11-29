package register

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/common/db"
	"github.com/kioku-project/kioku/pkg/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	db  *gorm.DB
	app *fiber.App
}

func New() (s *LoginService) {
	_ = godotenv.Load(".env", ".env.example")
	db.InitializeDB()

	app := fiber.New()
	s = &LoginService{db: db.InitializeDB(), app: app}
	app.Post("/api/register", s.Register)
	return
}

func (s *LoginService) Listen(addr string) error {
	return s.app.Listen(addr)
}

func (s *LoginService) Shutdown() error {
	return s.app.Shutdown()
}

func (s *LoginService) Register(ctx *fiber.Ctx) error {

	data := map[string]string{}
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	newUser := user.User{
		Email: data["email"],
	}
	if err := s.db.Where(&newUser).First(&newUser).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusBadRequest, "This email is already registered")
	}
	if data["name"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No Name given")
	}
	newUser.Name = data["name"]

	// TODO: check password requirements
	if data["password"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No Password given")
	}

	// encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.MinCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error while hashing password")
	}
	newUser.Password = string(hash)
	s.db.Create(&newUser)

	return ctx.SendString("User registered successfully")
}
