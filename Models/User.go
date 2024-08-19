package Models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint `gorm:"primaryKey"`
	Username    string
	PhoneNumber string
	Password    string
	FullName    string
	Email       string
	Age         uint8
	Birthday    time.Time
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
