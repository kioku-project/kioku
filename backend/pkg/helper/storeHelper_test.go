package helper_test

import (
	"testing"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func noErrorStoreFunc(string) (string, error) {
	return "", nil
}

func errorStoreFunc(string) (string, error) {
	return "", helper.ErrStoreNoEntryWithID
}

func TestStoreHelper(t *testing.T) {
	// NoError
	entity, err := helper.FindStoreEntity[string](
		noErrorStoreFunc,
		id,
		id,
	)
	assert.Equal(t, "", entity)
	assert.NoError(t, err)

	// Error
	entity, err = helper.FindStoreEntity[string](
		errorStoreFunc,
		id,
		id,
	)
	assert.Equal(t, "", entity)
	assert.Error(t, err)
}