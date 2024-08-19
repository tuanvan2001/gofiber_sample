package Models

import (
	"time"
)

type Event struct {
	ID        uint   `gorm:"primaryKey"`
	EventName string `gorm:"not null"`
	UserID    int    `gorm:"not null"`
	CreatedAt time.Time
}
