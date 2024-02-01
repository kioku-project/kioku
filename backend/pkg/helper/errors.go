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
	NotificationServiceID  ClientID = "services.notification"
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

func NewMicroCantLeaveAsLastAdminErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "You can't leave when you are the last admin")
}

func NewMicroUserAdmissionInProgressErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "user already invited")
}

func NewMicroCardSideNotInGivenCardErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "card side not in given card")
}

func NewMicroDeckTypeNotValidErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid deck type")
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

func NewMicroInvalidPasswordErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "invalid password")
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

func NewMicroDeckAlreadyFavoriteErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "This Deck is already your favorite")
}

func NewFiberBadRequestErr(detail string) error {
	return fiber.NewError(fiber.StatusBadRequest, detail)
}

func NewFiberMissingEmailErr() error {
	return fiber.NewError(fiber.StatusBadRequest, "no Email provided")
}

func NewFiberMissingNameErr() error {
	return fiber.NewError(fiber.StatusBadRequest, "no Name provided")
}

func NewFiberMissingPasswordErr() error {
	return fiber.NewError(fiber.StatusBadRequest, "no Password provided")
}

func NewFiberMissingDeckIDErr() error {
	return fiber.NewError(fiber.StatusBadRequest, "no deckID provided")
}

func NewFiberUnauthorizedErr(detail string) error {
	return fiber.NewError(fiber.StatusUnauthorized, detail)
}
