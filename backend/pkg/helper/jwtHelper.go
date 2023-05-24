package helper

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	pem2 "encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-micro.dev/v4/logger"
)

func getJWTPrivateKey() (*ecdsa.PrivateKey, error) {

	pem, _ := pem2.Decode([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if pem == nil {
		panic("JWT Private Key not given")
	}
	priv, err := x509.ParseECPrivateKey(pem.Bytes)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	return priv, nil
}

func GetJWTPublicKey() (crypto.PublicKey, error) {

	priv, err := getJWTPrivateKey()
	if err != nil {
		return nil, err
	}
	return priv.Public(), nil
}

func ParseJWTToken(tokenString string) (*jwt.Token, error) {

	if tokenString == "" {
		return nil, errors.New("please re-authenticate")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		pub, err := GetJWTPublicKey()
		if err != nil {
			return nil, errors.New("could not parse JWT public / private keypair")
		}
		return pub, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CreateJWTTokenString(exp time.Time, id interface{}, email interface{}, name interface{}) (string, error) {

	priv, err := getJWTPrivateKey()
	if err != nil {
		return "", err
	}

	tokenClaims := jwt.MapClaims{
		"sub":   id,
		"email": email,
		"name":  name,
		"exp":   exp.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES512, tokenClaims)
	tokenString, err := token.SignedString(priv)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUserIDFromContext(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}
