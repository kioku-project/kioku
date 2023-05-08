package model

type RoleType string

const (
	RoleAdmin RoleType = "admin"
	RoleWrite RoleType = "write"
	RoleRead  RoleType = "read"
)

type GroupUserRole struct {
	GroupID  uint `gorm:"primaryKey"`
	UserID   uint `gorm:"primaryKey"`
	RoleType RoleType

	Group Group `gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User  User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
