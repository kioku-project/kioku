package helper

import (
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"go-micro.dev/v4/logger"
	"regexp"
)

type RegexPattern string

const (
	UserNameRegex         RegexPattern = "^[a-zA-Z0-9-._~]{3,20}$"
	GroupAndDeckNameRegex RegexPattern = "^[a-zA-Z0-9-._~ ]{3,20}$"
)

func IsAuthorized(groupRole pbCollaboration.GroupRole, requiredRole pbCollaboration.GroupRole) bool {
	return groupRole.Number() >= requiredRole.Number()
}

func CheckForValidName(name string, pattern RegexPattern, clientID ClientID) error {
	logger.Infof("Check name %s for valid format", name)
	if isMatch, err := regexp.MatchString(string(pattern), name); !isMatch {
		return NewMicroInvalidUserNameFormatErr(clientID)
	} else if err != nil {
		return err
	}
	logger.Infof("Name %s is valid", name)
	return nil
}
