package helper_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/stretchr/testify/assert"
	microErrors "go-micro.dev/v4/errors"
)

const (
	id = "id"
)

func TestErrors(t *testing.T) {
	microNotAuthorizedErr := helper.NewMicroNotAuthorizedErr(id)
	microNoEntryWithIDErr := helper.NewMicroNoEntryWithIDErr(id)
	microNoExistingUserWithEmailErr := helper.NewMicroNoExistingUserWithEmailErr(id)
	microUserAlreadyExistsErr := helper.NewMicroUserAlreadyExistsErr(id)
	microInvalidUserNameFormatErr := helper.NewMicroInvalidNameFormatErr(id)
	microInvalidParameterDataErr := helper.NewMicroInvalidParameterDataErr(id)
	microUserAlreadyInGroupErr := helper.NewMicroUserAlreadyInGroupErr(id)
	microUserAdmissionInProgressErr := helper.NewMicroUserAdmissionInProgressErr(id)
	microCardSideNotInGivenCardErr := helper.NewMicroCardSideNotInGivenCardErr(id)
	microHashingFailedErr := helper.NewMicroHashingFailedErr(id)
	microInvalidEmailOrPasswordErr := helper.NewMicroInvalidEmailOrPasswordErr(id)
	microNotSuccessfulResponseErr := helper.NewMicroNotSuccessfulResponseErr(id)
	fiberBadRequestErr := helper.NewFiberBadRequestErr(id)
	fiberUnauthorizedErr := helper.NewFiberUnauthorizedErr(id)

	assert.IsType(t, microErrors.Unauthorized(string(id), "user not authorized"), microNotAuthorizedErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "no entry with id"), microNoEntryWithIDErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "no existing user with email"), microNoExistingUserWithEmailErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "this user already exists"), microUserAlreadyExistsErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "invalid user name format"), microInvalidUserNameFormatErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "invalid parameter data"), microInvalidParameterDataErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "user already in group"), microUserAlreadyInGroupErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "user already invited"), microUserAdmissionInProgressErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "card side not in given card"), microCardSideNotInGivenCardErr)
	assert.IsType(t, microErrors.InternalServerError(string(id), "error while hashing password"), microHashingFailedErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "invalid email or password"), microInvalidEmailOrPasswordErr)
	assert.IsType(t, microErrors.BadRequest(string(id), "operation not successful"), microNotSuccessfulResponseErr)
	assert.IsType(t, fiber.NewError(fiber.StatusBadRequest, id), fiberBadRequestErr)
	assert.IsType(t, fiber.NewError(fiber.StatusUnauthorized, id), fiberUnauthorizedErr)
}
