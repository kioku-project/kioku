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
	FrontendServiceID      ClientID = "services.frontend"
	UserServiceID          ClientID = "services.user"
	SrsServiceID           ClientID = "services.srs"
)

var (
	ErrStoreRetryCountExceeded        = errors.New("exceeded retry count")
	ErrStoreNoEntryWithID             = errors.New("no entry with id")
	ErrStoreNoExistingUserWithEmail   = errors.New("no existing user with email")
	ErrStoreInvalidGroupRoleForChange = errors.New("user has invalid role for change")
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

func NewMicroInvalidParameterDataErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid parameter data")
}

func NewMicroUserAlreadyInGroupErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "user already in group")
}

func NewMicroCantLeaveDefaultGroupErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "You can't leave your default group")
}

func NewMicroUserAdmissionInProgressErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "user already invited")
}

func NewMicroCardSideNotInGivenCardErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "card side not in given card")
}

func NewMicroAlreadyRequestedErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "user access already requested")
}

func NewMicroAlreadyInvitedErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "user already invited")
}

func NewMicroHashingFailedErr(id ClientID) error {
	return microErrors.InternalServerError(string(id), "error while hashing password")
}

func NewMicroInvalidEmailOrPasswordErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid email or password")
}

func NewMicroNotSuccessfulResponseErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "operation not successful")
}

func NewMicroWrongRatingErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid rating")
}

func NewMicroWrongDeckIDErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "wrong deck id")
}

func NewFiberBadRequestErr(detail string) error {
	return fiber.NewError(fiber.StatusBadRequest, detail)
}

func NewFiberUnauthorizedErr(detail string) error {
	return fiber.NewError(fiber.StatusUnauthorized, detail)
}
