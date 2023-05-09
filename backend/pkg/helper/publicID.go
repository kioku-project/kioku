package helper

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	random  = *rand.New(rand.NewSource(time.Now().UnixNano()))
	charset = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func GeneratePublicID() string {
	randomPublicID := make([]byte, 6)
	for i := range randomPublicID {
		randomPublicID[i] = charset[random.Intn(len(charset))]
	}
	return string(randomPublicID)
}

func ConvertRandomIDToModelID(identifier byte, randomPublicID string) string {
	return fmt.Sprintf("#%c-%s", identifier, randomPublicID)
}
