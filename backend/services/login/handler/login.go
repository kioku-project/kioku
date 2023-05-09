package handler

import (
	"context"
	"errors"
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/kioku-project/kioku/services/login/proto"
	"github.com/kioku-project/kioku/store"
)

type Login struct{ store store.Store }

func New(s store.Store) *Login { return &Login{store: s} }

func (e *Login) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	logger.Infof("Received Loginservice.Login request: %v", req)
	user, err := e.store.FindUserByEmail(req.Email)
	if err != nil {
		return errors.New("This user does not exist")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.New("This email or password is wrong")
	}
	rsp.Name = user.Name
	rsp.Id = uint64(user.ID)
	logger.Infof("Name: %v", user.Name)
	return nil
}
