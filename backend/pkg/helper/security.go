package helper

import (
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	"go-micro.dev/v4/logger"
	"regexp"
)

var (
	UserNameRegex         = regexp.MustCompile("^[a-zA-Z0-9-._~]{3,20}$")
	GroupAndDeckNameRegex = regexp.MustCompile("^[a-zA-Z0-9-._~ ]{3,20}$")
)

func IsAuthorized(groupRole pbCommon.GroupRole, requiredRole pbCommon.GroupRole) bool {
	return groupRole.Number() >= requiredRole.Number()
}

func CheckForValidName(name string, pattern *regexp.Regexp, clientID ClientID) error {
	logger.Infof("Check name %s for valid format", name)
	if !pattern.MatchString(name) {
		return NewMicroInvalidNameFormatErr(clientID)
	}
	logger.Infof("Name %s is valid", name)
	return nil
}

func CheckForValidPassword(password string, clientID ClientID) error {
	logger.Infof("Check password for valid format")
	if len(password) < 3 {
		return NewMicroInvalidParameterDataErr(clientID)
	}
	logger.Infof("Password is valid")
	return nil
}
