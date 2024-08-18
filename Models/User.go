package Models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey"`
	uuid        uuid.UUID `gorm:"index"`
	Username    string
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
