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

func NewMicroNotAuthorizedErr(id ClientID) error {
	return microErrors.Unauthorized(string(id), "user not authorized")
}

func NewMicroNoEntryWithIDErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "no entry with id")
}

func NewMicroNoExistingUserWithEmailErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "no existing user with email")
}

func NewMicroUserAlreadyExistsErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "this user already exists")
}

func NewMicroInvalidUserNameFormatErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid user name format")
}

func NewMicroInvalidParameterData(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid parameter data")
}

func NewMicroUserAlreadyInGroupErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "user already in group")
}

func NewMicroUserAdmissionInProgress(id ClientID) error {
	return microErrors.BadRequest(string(id), "user already invited")
}

func NewMicroHashingFailedErr(id ClientID) error {
	return microErrors.InternalServerError(string(id), "error while hashing password")
}

func NewMicroInvalidEmailOrPasswordErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid email or password")
}

func NewFiberBadRequestErr(detail string) error {
	return fiber.NewError(fiber.StatusBadRequest, detail)
}

func NewFiberUnauthorizedErr(detail string) error {
	return fiber.NewError(fiber.StatusUnauthorized, detail)
}
