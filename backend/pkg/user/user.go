package user

import (
	"github.com/kioku-project/kioku/pkg/group"
)

type User struct {
	ID     uint
	Name   string
	EMail  string
	Groups []group.Group `gorm:"many2many:group_users"`
}
