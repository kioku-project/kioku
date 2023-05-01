package handler

import (
	"context"
	"errors"
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/kioku-project/kioku/services/register/proto"
	pblogin "github.com/kioku-project/kioku/services/login/proto"
	"github.com/kioku-project/kioku/store"
	"github.com/kioku-project/kioku/pkg/model"
)

type Register struct{
	store store.Store
	loginService pblogin.LoginService
}

func New(s store.Store, lS pblogin.LoginService) *Register { return &Register{store: s, loginService: lS} }

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

	dberr := e.store.RegisterNewUser(&newUser)
	if dberr != nil {
		logger.Infof("Error while inserting into db: %v", dberr.Error())
	}

	// Call the login service
	rspLogin, err := e.loginService.Login(context.TODO(), &pblogin.LoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		logger.Infof("Login service error: %v", err.Error())
		return err
	}

	// Print response of login service
	logger.Infof("Login service response: %v", rspLogin)

	rsp.Name = rspLogin.Name
	return nil
}
