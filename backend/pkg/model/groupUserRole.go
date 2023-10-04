package model

type RoleType string

const (
	RoleAdmin     RoleType = "admin"
	RoleWrite     RoleType = "write"
	RoleRead      RoleType = "read"
	RoleInvited   RoleType = "invited"
	RoleRequested RoleType = "requested"
)

type GroupUserRole struct {
	GroupID  string `gorm:"primaryKey"`
	UserID   string `gorm:"primaryKey"`
	RoleType RoleType
	Group    Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User     User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
