package helper

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

var ErrRetryCountExceeded = errors.New("exceeded retry count")

type PublicID struct {
	prefix rune
	id     string
}

func GeneratePublicID(prefix rune) PublicID {
	randomPublicID := make([]byte, 6)
	for i := range randomPublicID {
		randomPublicID[i] = charset[random.Intn(len(charset))]
	}
	return PublicID{prefix: prefix, id: string(randomPublicID)}
}

func (i PublicID) GetStringRepresentation() string {
	return fmt.Sprintf("%c-%s", i.prefix, i.id)
}

func FindFreePublicID[T, C any](
	db *gorm.DB,
	retries int,
	where func() (C, *T),
) (res C, err error) {
	var currentTry int
	var t T
	for {
		currentTry++
		if currentTry > retries {
			err = ErrRetryCountExceeded
			return
		}
		candidate, val := where()
		if err = db.Where(val).First(&t).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return candidate, nil
			}
			return
		}
	}
}
