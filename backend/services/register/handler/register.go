package handler

import (
	"context"
	"errors"
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/kioku-project/kioku/services/register/proto"
	"github.com/kioku-project/kioku/store"
	"github.com/kioku-project/kioku/pkg/model"
)

type Register struct{store store.Store}

func New(s store.Store) *Register { return &Register{store: s} }

func (e *Register) Register(ctx context.Context, req *pb.RegisterRequest, rsp *pb.RegisterResponse) error {
	logger.Infof("Received Register.Register request: %v", req)
	_, err := e.store.FindUserByEmail(req.Email)
	if err == nil {
		return errors.New("This user already exists")
	}

	// TODO: further username verification required (length, bad chars, ...)
	newUser := model.User{
		Email: req.Email,
		Name: req.Name,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return errors.New("Error while hashing password")
	}
	newUser.Password = string(hash)
	e.store.RegisterNewUser(&newUser)

	rsp.Name = newUser.Name
	return nil
}
