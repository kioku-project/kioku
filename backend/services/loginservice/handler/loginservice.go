package handler

import (
	"context"
	"errors"
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/kioku-project/kioku/services/loginservice/proto"
	"github.com/kioku-project/kioku/store"
)

type Loginservice struct{store store.Store}

func New(s store.Store) *Loginservice { return &Loginservice{store: s} }

func (e *Loginservice) Login(ctx context.Context, req *pb.LoginserviceRequest, rsp *pb.LoginserviceResponse) error {
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
	return nil
}
