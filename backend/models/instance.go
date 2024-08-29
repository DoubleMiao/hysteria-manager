package models

import (
    "time"
)

type Instance struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"unique;not null"`
    Port      int       `gorm:"not null"`
    CreatedAt time.Time
    ExpiresAt time.Time
}
