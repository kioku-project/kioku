package model

type User struct {
	ID       uint    `gorm:"primaryKey;" json:"id"`
	Name     string  `gorm:"not null;" json:"name"`
	Email    string  `gorm:"unique;not null;" json:"email"`
	Password string  `gorm:"not null;" json:"password"`
	Groups   []Group `gorm:"many2many:group_user_roles;" json:"groups"`
}
