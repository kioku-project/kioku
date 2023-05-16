package helper

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

var (
	random  = *rand.New(rand.NewSource(time.Now().UnixNano()))
	charset = "abcdefghijklmnopqrstuvwxyz0123456789"
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

func FindFreePublicID[T any](
	db *gorm.DB,
	retries int,
	generator func(prefix rune) PublicID,
	prefix rune,
	where func(candidate string) *T,
) (res string, err error) {
	var currentTry int
	var t T
	for {
		currentTry++
		if currentTry > retries {
			err = errors.New("exceeded retry count")
			return
		}
		candidate := generator(prefix)
		if err = db.Where(where(candidate.GetStringRepresentation())).First(&t).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return candidate.GetStringRepresentation(), nil
			}
			return
		}
	}
}
