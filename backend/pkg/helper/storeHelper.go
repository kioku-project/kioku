package helper

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
)

func FindStoreEntity[C any](
	ctx context.Context,
	storeFunction func(context.Context, string) (C, error),
	ID string,
	callContext ClientID,
) (entity C, err error) {
	if entity, err = storeFunction(ctx, ID); errors.Is(err, ErrStoreNoEntryWithID) {
		err = NewMicroNoEntryWithIDErr(callContext)
		logger.Infof("Throwing error (%v) for query for id %s", err, ID)
	} else {
		logger.Infof("Found entity/-ies with/by id %s", ID)
	}
	return
}
