package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	microErrors "go-micro.dev/v4/errors"
)

type ClientID string

const (
	CardDeckServiceID      ClientID = "services.cardDeck"
	CollaborationServiceID ClientID = "services.collaboration"
	UserServiceID          ClientID = "services.user"
)

var (
	ErrStoreRetryCountExceeded      = errors.New("exceeded retry count")
	ErrStoreNoEntryWithID           = errors.New("no entry with id")
	ErrStoreNoExistingUserWithEmail = errors.New("no existing user with email")
)

func ErrMicroNotAuthorized(id ClientID) error {
	return microErrors.Unauthorized(string(id), "user not authorized")
}

func ErrMicroNoEntryWithID(id ClientID) error {
	return microErrors.BadRequest(string(id), "no entry with id")
}

func ErrMicroNoExistingUserWithEmail(id ClientID) error {
	return microErrors.BadRequest(string(id), "no existing user with email")
}

func ErrMicroUserAlreadyExists(id ClientID) error {
	return microErrors.BadRequest(string(id), "this user already exists")
}

func ErrMicroInvalidUserNameFormat(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid user name format")
}

func ErrMicroHashingFailed(id ClientID) error {
	return microErrors.InternalServerError(string(id), "error while hashing password")
}

func ErrMicroInvalidEmailOrPassword(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid email or password")
}

func ErrFiberBadRequest(detail string) error {
	return fiber.NewError(fiber.StatusBadRequest, detail)
}

func ErrFiberUnauthorized(detail string) error {
	return fiber.NewError(fiber.StatusUnauthorized, detail)
}

func ErrFiberInternalServerError(detail string) error {
	return fiber.NewError(fiber.StatusInternalServerError, detail)
}
