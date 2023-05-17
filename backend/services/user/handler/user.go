package handler

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"

	"github.com/kioku-project/kioku/pkg/model"
	pbcollab "github.com/kioku-project/kioku/services/collaboration/proto"
	pb "github.com/kioku-project/kioku/services/user/proto"
	"github.com/kioku-project/kioku/store"
)

type User struct {
	store                store.UserStore
	collaborationService pbcollab.CollaborationService
}

func New(s store.UserStore, cS pbcollab.CollaborationService) *User {
	return &User{store: s, collaborationService: cS}
}

func (e *User) Register(ctx context.Context, req *pb.RegisterRequest, rsp *pb.NameIDResponse) error {
	logger.Infof("Received User.Register request: %v", req)
	_, err := e.store.FindUserByEmail(req.Email)
	if err == nil {
		return errors.New("this user already exists")
	}

	// TODO: further username verification required (length, bad chars, ...)
	newUser := model.User{
		Email: req.Email,
		Name:  req.Name,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return errors.New("error while hashing password")
	}
	newUser.Password = string(hash)

	dberr := e.store.RegisterNewUser(&newUser)
	if dberr != nil {
		logger.Infof("Error while inserting into db: %v", dberr.Error())
	}

	rspGroup, err := e.collaborationService.CreateNewGroupWithAdmin(context.TODO(), &pbcollab.CreateGroupRequest{UserID: newUser.ID, GroupName: "Home Group"})
	if err != nil || !rspGroup.Success {
		logger.Infof("Collaboration service error: %v", err.Error())
		return err
	}

	rsp.Name = newUser.Name
	rsp.ID = newUser.ID
	logger.Infof("Name: %v", newUser.Name)
	return nil
}

func (e *User) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.NameIDResponse) error {
	logger.Infof("Received User.Login request: %v", req)
	user, err := e.store.FindUserByEmail(req.Email)
	if err != nil {
		return errors.New("this user does not exist")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.New("this email or password is wrong")
	}
	rsp.Name = user.Name
	rsp.ID = user.ID
	logger.Infof("Name: %v", user.Name)
	return nil
}

func (e *User) GetUserIDFromEmail(ctx context.Context, req *pb.UserIDRequest, rsp *pb.UserIDResponse) error {
	logger.Infof("Received User.GetUserIDFromEmail request: %v", req)
	user, err := e.store.FindUserByEmail(req.Email)
	if err != nil {
		return err
	}
	rsp.ID = user.ID
	return nil
}
