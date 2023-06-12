package helper_test

import (
	"testing"
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func TestJWTHelper(t *testing.T) {
	_, err := helper.ParseJWTToken("")
	assert.Error(t, err)

	// Panics because no JWT_PRIVATE_KEY is given
	assert.Panics(t, func() {
		helper.GetJWTPublicKey()
	})

	// Panics because no JWT_PRIVATE_KEY is given
	assert.Panics(t, func() {
		helper.CreateJWTTokenString(time.Now(), "id", "mail", "name")
	})
}