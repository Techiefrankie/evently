package models

import (
	"time"
)

type Event struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"unique;not null;size:100"`
	Description string    `gorm:"not null;size:255"`
	Location    string    `gorm:"not null;size:255"`
	DateTime    time.Time `gorm:"not null"`
	UserId      int       `gorm:"not null"`
}
