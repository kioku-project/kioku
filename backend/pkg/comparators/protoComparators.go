package comparators

import (
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	"strings"
)

func GroupUserProtoRoleComparator(a, b *pbCommon.User) int {
	if val := int(b.GroupRole.Number()) - int(a.GroupRole.Number()); val != 0 {
		return val
	}
	return strings.Compare(a.UserName, b.UserName)
}
