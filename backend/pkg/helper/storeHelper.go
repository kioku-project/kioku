package helper

import (
	"errors"
	"go-micro.dev/v4/logger"
)

func FindStoreEntity[C any](
	storeFunction func(string) (C, error),
	ID string,
	callContext ClientID,
) (entity C, err error) {
	entity, err = storeFunction(ID)
	if entity, err = storeFunction(ID); errors.Is(err, ErrStoreNoEntryWithID) {
		err = NewMicroNoEntryWithIDErr(callContext)
	}
	if err == nil {
		logger.Infof("Found entity/-ies with/by id %s", ID)
	} else {
		logger.Infof("Throwing error (%s) for query for id %s", err.Error(), ID)
	}
	return
}
