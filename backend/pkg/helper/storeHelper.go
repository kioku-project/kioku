package helper

import (
	"errors"
	"go-micro.dev/v4/logger"
)

func FindEntityWrapper[C any](
	storeFunction func(string) (C, error),
	ID string,
	callContext ClientID,
) (entity C, err error) {
	entity, err = storeFunction(ID)
	if err != nil {
		if errors.Is(err, ErrStoreNoEntryWithID) {
			err = ErrMicroNoEntryWithID(callContext)
		}
	}
	logger.Infof("Found entity/-ies with/by id %s", ID)
	return
}
