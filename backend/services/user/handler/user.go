package handler

import (
	"context"
	"errors"
	"net/mail"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/helper"
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"

	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/store"
)

type User struct {
	store                store.UserStore
	collaborationService pbCollaboration.CollaborationService
}

func New(s store.UserStore, cS pbCollaboration.CollaborationService) *User {
	return &User{store: s, collaborationService: cS}
}

func (e *User) Register(ctx context.Context, req *pbCommon.User, rsp *pbCommon.Success) error {
	logger.Infof("Received User.Register request: email: %v", req.UserEmail)
	if _, err := e.store.FindUserByEmail(ctx, req.UserEmail); err == nil {
		return helper.NewMicroUserAlreadyExistsErr(helper.UserServiceID)
	} else if !errors.Is(err, helper.ErrStoreNoExistingUserWithEmail) {
		return err
	}
	if err := helper.CheckForValidName(req.UserName, helper.UserNameRegex, helper.UserServiceID); err != nil {
		return err
	}
	addr, err := mail.ParseAddress(req.UserEmail)
	if err != nil {
		return err
	}
	newUser := model.User{
		Email: addr.Address,
		Name:  req.UserName,
	}
	if req.UserPassword == "" {
		return helper.NewMicroInvalidEmailOrPasswordErr(helper.UserServiceID)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.MinCost)
	if err != nil {
		return helper.NewMicroHashingFailedErr(helper.UserServiceID)
	}
	newUser.Password = string(hash)
	err = e.store.RegisterNewUser(ctx, &newUser)
	if err != nil {
		return err
	}
	_, err = e.collaborationService.CreateNewGroupWithAdmin(ctx, &pbCommon.GroupRequest{
		UserID: newUser.ID,
		Group: &pbCommon.Group{
			GroupName:        "Home Group",
			GroupDescription: "Your personal deck space",
			IsDefault:        true,
			GroupType:        pbCommon.GroupType_CLOSED,
		},
	})
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully created new user with id %s", newUser.ID)
	return nil
}

func (e *User) VerifyUserExists(ctx context.Context, req *pbCommon.User, rsp *pbCommon.Success) error {
	user, err := e.store.FindUserByEmail(ctx, req.UserEmail)
	if err != nil {
		return err
	}
	if user.ID == req.UserID {
		rsp.Success = true
	}
	return nil
}

func (e *User) DeleteUser(ctx context.Context, req *pbCommon.User, rsp *pbCommon.Success) error {
	logger.Infof("Received User.Delete request: %v", req)
	user, err := e.store.FindUserByID(ctx, req.UserID)
	if err != nil {
		logger.Errorf("Could not find user with userid: %s: %s", user.ID, err)
		return err
	}
	// Handle groups that the user is part of
	groupRes, err := e.collaborationService.GetUserGroups(ctx, &pbCommon.User{UserID: user.ID})
	if err != nil {
		return err
	}
	for _, group := range groupRes.Groups {
		_, err := e.collaborationService.LeaveGroup(ctx, &pbCommon.GroupRequest{
			UserID: req.UserID,
			Group: &pbCommon.Group{
				GroupID: group.GroupID,
			},
		})
		if err != nil {
			return err
		}
	}
	err = e.store.DeleteUser(ctx, user)
	if err != nil {
		logger.Errorf("An error occurred while trying to delete user with ID '%s': %s", user.ID, err)
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted user with id %s", user.ID)
	return nil
}

func (e *User) Login(ctx context.Context, req *pbCommon.User, rsp *pbCommon.User) error {
	logger.Infof("Received User.Login request: email: %v", req.UserEmail)
	user, err := e.store.FindUserByEmail(ctx, req.UserEmail)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoExistingUserWithEmail) {
			return helper.NewMicroNoExistingUserWithEmailErr(helper.UserServiceID)
		}
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.UserPassword))
	if err != nil {
		return helper.NewMicroInvalidEmailOrPasswordErr(helper.UserServiceID)
	}
	rsp.UserName = user.Name
	rsp.UserID = user.ID
	logger.Infof("Successfully logged in user with id %s", user.ID)
	return nil
}

func (e *User) GetUserIDFromEmail(ctx context.Context, req *pbCommon.User, rsp *pbCommon.User) error {
	logger.Infof("Received User.GetUserIDFromEmail request: %v", req)
	user, err := e.store.FindUserByEmail(ctx, req.UserEmail)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoExistingUserWithEmail) {
			return helper.NewMicroNoExistingUserWithEmailErr(helper.UserServiceID)
		}
		return err
	}
	rsp.UserID = user.ID
	logger.Infof("Found user with id %s", user.ID)
	return nil
}
func (e *User) GetUserInformation(
	ctx context.Context,
	req *pbCommon.Users,
	rsp *pbCommon.Users,
) error {
	logger.Infof("Received User.GetUserInformation request: %v", req)
	rsp.Users = make([]*pbCommon.User, len(req.Users))
	for i, user := range req.Users {
		user, err := e.store.FindUserByID(ctx, user.UserID)
		if err != nil {
			return err
		}
		rsp.Users[i] = &pbCommon.User{
			UserID:    user.ID,
			UserName:  user.Name,
			UserEmail: user.Email,
		}
	}
	logger.Infof("Found %d users for given IDs", len(rsp.Users))
	return nil
}

func (e *User) GetUserProfileInformation(
	ctx context.Context,
	req *pbCommon.User,
	rsp *pbCommon.User,
) error {
	logger.Infof("Received User.GetUserProfileInformation request: %v", req)
	user, err := e.store.FindUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}
	*rsp = *converter.StoreUserToProtoUserProfileInformationResponseConverter(*user)
	logger.Infof("Found profile information for user with id %s", req.UserID)
	return nil
}

func (e *User) ModifyUserProfileInformation(
	ctx context.Context,
	req *pbCommon.User,
	rsp *pbCommon.Success,
) error {
	logger.Infof("Received User.ModifyUserProfileInformation request: %v", req)
	user, err := e.store.FindUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}
	if req.UserName != "" {
		err = helper.CheckForValidName(req.UserName, helper.UserNameRegex, helper.UserServiceID)
		if err != nil {
			return err
		}
		user.Name = req.UserName
	}
	if req.UserEmail != "" {
		addr, err := mail.ParseAddress(req.UserEmail)
		if err != nil {
			return err
		}
		user.Email = addr.Address
	}
	if req.UserPassword != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.MinCost)
		if err != nil {
			return helper.NewMicroHashingFailedErr(helper.UserServiceID)
		}
		user.Password = string(hash)
	}
	if err = e.store.ModifyUser(ctx, user); err != nil {
		return err
	}
	logger.Infof("Modified profile information for user with id %s", req.UserID)
	rsp.Success = true
	return nil
}
