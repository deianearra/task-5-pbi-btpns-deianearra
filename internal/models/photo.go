package models

import (
	"time"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey;not null"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoURL  string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
