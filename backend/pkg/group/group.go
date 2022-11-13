package group

import (
	"github.com/kioku-project/kioku/pkg/groupUser"
)

type Group struct {
	ID    uint
	Name  string
	Users []groupUser.GroupUser `gorm:"many2many:group_users"`
}
