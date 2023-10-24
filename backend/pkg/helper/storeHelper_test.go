package helper_test

import (
	"context"
	"testing"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func noErrorStoreFunc(context.Context, string) (string, error) {
	return "", nil
}

func errorStoreFunc(context.Context, string) (string, error) {
	return "", helper.ErrStoreNoEntryWithID
}

func TestStoreHelper(t *testing.T) {
	// NoError
	entity, err := helper.FindStoreEntity[string](
		context.TODO(),
		noErrorStoreFunc,
		id,
		id,
	)
	assert.Equal(t, "", entity)
	assert.NoError(t, err)

	// Error
	entity, err = helper.FindStoreEntity[string](
		context.TODO(),
		errorStoreFunc,
		id,
		id,
	)
	assert.Equal(t, "", entity)
	assert.Error(t, err)
}