package model

type Group struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Users []User `gorm:"many2many:group_user_roles;"`
}
