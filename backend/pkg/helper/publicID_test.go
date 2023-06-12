package helper_test

import (
	"testing"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/stretchr/testify/assert"
)

const (
	prefixRune = 'S'
	prefixString = "S"
)

func TestPublicID(t *testing.T) {
	publicID := helper.GenerateID(prefixRune)
	assert.Equal(t, prefixString, string(publicID.GetStringRepresentation()[0]))
}