package store

import "github.com/kioku-project/kioku/pkg/model"

type Store interface {
	FindUserByEmail(email string) (*model.User, error)
	RegisterNewUser(newUser *model.User) (error)
}
