package model

type User struct {
	ID       uint          `json:"id"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Groups   []Group `gorm:"many2many:group_users" json:"groups"`
}
