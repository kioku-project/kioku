package login

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	// TODO Restrict CORS
	app.Use(cors.New())
	s = &LoginService{db: db.InitializeDB(), app: app}
	app.Post("/api/login", s.Login)
	return
}

func (s *LoginService) Listen(addr string) error {
	return s.app.Listen(addr)
}

func (s *LoginService) Shutdown() error {
	return s.app.Shutdown()
}

func (s *LoginService) Login(ctx *fiber.Ctx) error {

	data := map[string]string{}
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["email"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No E-Mail given")
	}

	if data["password"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No Password given")
	}

	newUser := user.User{
		Email: data["email"],
	}

	if err := s.db.Where(&newUser).First(&newUser).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(data["password"]))
		if err == nil {
			return ctx.SendString(newUser.Name)
		}
	}
	return fiber.NewError(fiber.StatusBadRequest, "This email or password is wrong")
}
