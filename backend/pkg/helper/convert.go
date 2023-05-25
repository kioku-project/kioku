package helper

import pbUser "github.com/kioku-project/kioku/services/user/proto"

func ConvertToProtoUserIDs[C any](toConvert []C, convert func(C) string) []*pbUser.UserID {
	userIDs := make([]*pbUser.UserID, len(toConvert))
	for i, user := range toConvert {
		userIDs[i] = &pbUser.UserID{
			UserID: convert(user),
		}
	}
	return userIDs
}
