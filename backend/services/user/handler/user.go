package handler

import (
	"context"
	"errors"
	"github.com/kioku-project/kioku/pkg/helper"
	"regexp"

	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"

	"github.com/kioku-project/kioku/pkg/model"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pb "github.com/kioku-project/kioku/services/user/proto"
	"github.com/kioku-project/kioku/store"
)

type User struct {
	store                store.UserStore
	collaborationService pbCollaboration.CollaborationService
}

func New(s store.UserStore, cS pbCollaboration.CollaborationService) *User {
	return &User{store: s, collaborationService: cS}
}

func (e *User) Register(ctx context.Context, req *pb.RegisterRequest, rsp *pb.NameIDResponse) error {
	logger.Infof("Received User.Register request: %v", req)
	if _, err := e.store.FindUserByEmail(req.Email); err == nil {
		return helper.NewMicroUserAlreadyExistsErr(helper.UserServiceID)
	} else if !errors.Is(err, helper.ErrStoreNoExistingUserWithEmail) {
		return err
	}
	if isMatch, err := regexp.MatchString("^[a-zA-Z0-9-._~]{3,20}$", req.Name); !isMatch {
		return helper.NewMicroInvalidUserNameFormatErr(helper.UserServiceID)
	} else if err != nil {
		return err
	}
	newUser := model.User{
		Email: req.Email,
		Name:  req.Name,
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return helper.NewMicroHashingFailedErr(helper.UserServiceID)
	}
	newUser.Password = string(hash)
	err = e.store.RegisterNewUser(&newUser)
	if err != nil {
		return err
	}
	_, err = e.collaborationService.CreateNewGroupWithAdmin(ctx, &pbCollaboration.CreateGroupRequest{UserID: newUser.ID, GroupName: "Home Group", IsDefault: true})
	if err != nil {
		return err
	}
	rsp.Name = newUser.Name
	rsp.ID = newUser.ID
	logger.Infof("Successfully created new user with id %s", newUser.ID)
	return nil
}

func (e *User) Login(_ context.Context, req *pb.LoginRequest, rsp *pb.NameIDResponse) error {
	logger.Infof("Received User.Login request: %v", req)
	user, err := e.store.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoExistingUserWithEmail) {
			return helper.NewMicroNoExistingUserWithEmailErr(helper.UserServiceID)
		}
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return helper.NewMicroInvalidEmailOrPasswordErr(helper.UserServiceID)
	}
	rsp.Name = user.Name
	rsp.ID = user.ID
	logger.Infof("Successfully logged in user with id %s", user.ID)
	return nil
}

func (e *User) GetUserIDFromEmail(_ context.Context, req *pb.UserIDRequest, rsp *pb.UserID) error {
	logger.Infof("Received User.GetUserIDFromEmail request: %v", req)
	user, err := e.store.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.NewMicroNoEntryWithIDErr(helper.UserServiceID)
		}
		return err
	}
	rsp.UserID = user.ID
	logger.Infof("Found user with id %s", user.ID)
	return nil
}

func (e *User) GetUserInformation(_ context.Context, req *pb.UserInformationRequest, rsp *pb.UserInformationResponse) error {
	logger.Infof("Received User.GetUserInformation request: %v", req)
	rsp.Users = make([]*pb.UserInformation, len(req.UserIDs))
	for i, user := range req.UserIDs {
		user, err := e.store.FindUserByID(user.UserID)
		if err != nil {
			return err
		}
		rsp.Users[i] = &pb.UserInformation{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
		}
	}
	logger.Infof("Found %d users for given IDs", len(rsp.Users))
	return nil
}
