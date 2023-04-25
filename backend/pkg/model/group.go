package model

type Group struct {
	ID    uint
	Name  string
	Users []GroupUser `gorm:"many2many:group_users"`
}
