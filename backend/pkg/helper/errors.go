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
	return microErrors.Unauthorized(string(id), "User not authorized")
}

func NewMicroNoEntryWithIDErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "No entry with id")
}

func NewMicroNoExistingUserWithEmailErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "User does not exist")
}

func NewMicroUserAlreadyExistsErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "This user already exists")
}

func NewMicroInvalidNameFormatErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Invalid name format")
}

func NewMicroInvalidParameterDataErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Invalid parameter data")
}

func NewMicroMissingParameterDataErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Missing request parameters")
}

func NewMicroUserAlreadyInGroupErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "User already in group")
}

func NewMicroCantLeaveDefaultGroupErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "You cannot leave your default group")
}

func NewMicroCantLeaveAsLastAdminErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "You cannot leave as last admin")
}

func NewMicroCantInviteToHomegroupErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "You cannot invite users to your homegroup")
}

func NewMicroUserAdmissionInProgressErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "User already invited")
}

func NewMicroCardSideNotInGivenCardErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Card side not in given card")
}

func NewMicroDeckTypeNotValidErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Invalid deck type")
}

func NewMicroAlreadyRequestedErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "User access already requested")
}

func NewMicroAlreadyInvitedErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "User already invited")
}

func NewMicroCantModifyGroupAdminErr(id ClientID) error {
	return microErrors.Unauthorized(string(id), "You cannot modify group admins")
}

func NewMicroCantKickGroupAdminErr(id ClientID) error {
	return microErrors.Unauthorized(string(id), "You cannot kick group admins")
}

func NewMicroHashingFailedErr(id ClientID) error {
	return microErrors.InternalServerError(string(id), "Error while hashing password")
}

func NewMicroInvalidEmailOrPasswordErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Invalid email or password")
}

func NewMicroNotSuccessfulResponseErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Operation not successful")
}

func NewMicroWrongRatingErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Invalid rating")
}

func NewMicroWrongDeckIDErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "Wrong deck id")
}

func NewMicroDeckAlreadyFavoriteErr(id ClientID) error {
	return microErrors.BadRequest(string(id), "This Deck is already your favorite")
}

func NewFiberReauthenticateError() error {
	return fiber.NewError(fiber.StatusUnauthorized, "Please re-authenticate")
}

func NewFiberBadRequestErr(detail string) error {
	return fiber.NewError(fiber.StatusBadRequest, detail)
}

func NewFiberMissingDeckIDErr() error {
	return fiber.NewError(fiber.StatusBadRequest, "No deckID provided")
}

func NewFiberUnauthorizedErr(detail string) error {
	return fiber.NewError(fiber.StatusUnauthorized, detail)
}
