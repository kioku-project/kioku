package helper

import pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"

func IsAuthorized(groupRole pbCollaboration.GroupRole, requiredRole pbCollaboration.GroupRole) bool {
	if groupRole.Number() >= requiredRole.Number() {
		return true
	}
	return false
}
