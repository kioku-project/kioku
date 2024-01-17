package model

type PushSubscription struct {
	ID       string `gorm:"primaryKey"`
	UserID   string `gorm:"not null"`
	User     User
	Endpoint string `gorm:"not null"`
	P256DH   string `gorm:"not null"`
	Auth     string `gorm:"not null"`
}
