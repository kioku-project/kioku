package helper

import (
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"go-micro.dev/v4/logger"
	"regexp"
)

var (
	UserNameRegex         = regexp.MustCompile("^[a-zA-Z0-9-._~]{3,20}$")
	GroupAndDeckNameRegex = regexp.MustCompile("^[a-zA-Z0-9-._~ ]{3,20}$")
)

func IsAuthorized(groupRole pbCollaboration.GroupRole, requiredRole pbCollaboration.GroupRole) bool {
	return groupRole.Number() >= requiredRole.Number()
}

func CheckForValidName(name string, pattern *regexp.Regexp, clientID ClientID) error {
	logger.Infof("Check name %s for valid format", name)
	if !pattern.MatchString(name) {
		return NewMicroInvalidUserNameFormatErr(clientID)
	}
	logger.Infof("Name %s is valid", name)
	return nil
}
