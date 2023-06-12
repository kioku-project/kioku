package helper_test

import (
	"testing"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func TestSecurityTest(t *testing.T) {
	auth := helper.IsAuthorized(3, 2)
	assert.Equal(t, true, auth)

	// Valid name
	err := helper.CheckForValidName("validname", helper.UserNameRegex, "dummy")
	assert.NoError(t, err)

	// Invalid name
	err = helper.CheckForValidName("", helper.UserNameRegex, "dummy")
	assert.Error(t, err)
}