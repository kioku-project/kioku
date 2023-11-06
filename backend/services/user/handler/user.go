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
	logger.Infof("Received User.Register request: email: %v", req.UserEmail)
	if _, err := e.store.FindUserByEmail(req.UserEmail); err == nil {
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
	err = e.store.RegisterNewUser(&newUser)
	if err != nil {
		return err
	}
	_, err = e.collaborationService.CreateNewGroupWithAdmin(ctx, &pbCollaboration.CreateGroupRequest{
		UserID:           newUser.ID,
		GroupName:        "Home Group",
		GroupDescription: "Your personal deck space",
		IsDefault:        true,
	})
	if err != nil {
		return err
	}
	rsp.UserName = newUser.Name
	rsp.UserID = newUser.ID
	logger.Infof("Successfully created new user with id %s", newUser.ID)
	return nil
}

func (e *User) VerifyUserExists(_ context.Context, req *pb.VerificationRequest, rsp *pb.SuccessResponse) error {
	user, err := e.store.FindUserByEmail(req.UserEmail)
	if err != nil {
		return err
	}
	if user.ID == req.UserID {
		rsp.Success = true
	}
	return nil
}

func (e *User) DeleteUser(ctx context.Context, req *pb.UserID, rsp *pb.SuccessResponse) error {
	logger.Infof("Received User.Delete request: %v", req)
	user, err := e.store.FindUserByID(req.UserID)
	if err != nil {
		logger.Errorf("Could not find user with userid: %s: %s", user.ID, err)
		return err
	}
	// Handle groups that the user is part of
	groupRes, err := e.collaborationService.GetUserGroups(ctx, &pbCollaboration.UserIDRequest{UserID: user.ID})
	if err != nil {
		return err
	}
	for _, group := range groupRes.Groups {
		_, err := e.collaborationService.LeaveGroup(ctx, &pbCollaboration.GroupRequest{
			UserID:  req.UserID,
			GroupID: group.Group.GroupID,
		})
		if err != nil {
			return err
		}
	}
	err = e.store.DeleteUser(user)
	if err != nil {
		logger.Errorf("An error occurred while trying to delete user with ID '%s': %s", user.ID, err)
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted user with id %s", user.ID)
	return nil
}

func (e *User) Login(_ context.Context, req *pb.LoginRequest, rsp *pb.NameIDResponse) error {
	logger.Infof("Received User.Login request: email: %v", req.UserEmail)
	user, err := e.store.FindUserByEmail(req.UserEmail)
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

func (e *User) GetUserIDFromEmail(_ context.Context, req *pb.UserIDRequest, rsp *pb.UserID) error {
	logger.Infof("Received User.GetUserIDFromEmail request: %v", req)
	user, err := e.store.FindUserByEmail(req.UserEmail)
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
	_ context.Context,
	req *pb.UserInformationRequest,
	rsp *pb.UserInformationResponse,
) error {
	logger.Infof("Received User.GetUserInformation request: %v", req)
	rsp.Users = make([]*pb.UserInformation, len(req.UserIDs))
	for i, user := range req.UserIDs {
		user, err := e.store.FindUserByID(user.UserID)
		if err != nil {
			return err
		}
		rsp.Users[i] = &pb.UserInformation{
			UserID:    user.ID,
			UserName:  user.Name,
			UserEmail: user.Email,
		}
	}
	logger.Infof("Found %d users for given IDs", len(rsp.Users))
	return nil
}

func (e *User) GetUserProfileInformation(
	_ context.Context,
	req *pb.UserID,
	rsp *pb.UserProfileInformationResponse,
) error {
	logger.Infof("Received User.GetUserProfileInformation request: %v", req)
	user, err := e.store.FindUserByID(req.UserID)
	if err != nil {
		return err
	}
	*rsp = *converter.StoreUserToProtoUserProfileInformationResponseConverter(*user)
	logger.Infof("Found profile information for user with id %s", req.UserID)
	return nil
}

func (e *User) ModifyUserProfileInformation(
	_ context.Context,
	req *pb.ModifyRequest,
	rsp *pb.SuccessResponse,
) error {
	logger.Infof("Received User.ModifyUserProfileInformation request: %v", req)
	user, err := e.store.FindUserByID(req.UserID)
	if err != nil {
		return err
	}
	if req.UserName != nil {
		err = helper.CheckForValidName(*req.UserName, helper.UserNameRegex, helper.UserServiceID)
		if err != nil {
			return err
		}
		user.Name = *req.UserName
	}
	if req.UserEmail != nil {
		addr, err := mail.ParseAddress(*req.UserEmail)
		if err != nil {
			return err
		}
		user.Email = addr.Address
	}
	if req.UserPassword != nil {
		if req.UserPassword == "" {
			return helper.NewMicroInvalidEmailOrPasswordErr(helper.UserServiceID)
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(*req.UserPassword), bcrypt.MinCost)
		if err != nil {
			return helper.NewMicroHashingFailedErr(helper.UserServiceID)
		}
		user.Password = string(hash)
	}
	if err = e.store.ModifyUser(user); err != nil {
		return err
	}
	logger.Infof("Modified profile information for user with id %s", req.UserID)
	rsp.Success = true
	return nil
}
