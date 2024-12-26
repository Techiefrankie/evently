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
	UserId      int       `gorm:"not null"`                        // Foreign key field
	User        User      `gorm:"foreignKey:UserId;references:Id"` // Foreign key relationship
}

type User struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"not null;size:100"`
	LastName  string `gorm:"not null;size:100"`
	Email     string `gorm:"unique;not null;size:100"`
	Password  string `gorm:"not null;size:255"`
}

type Registration struct {
	Id               uint      `gorm:"primaryKey;autoIncrement"`
	UserId           int       `gorm:"not null"`                         // Foreign key field
	User             User      `gorm:"foreignKey:UserId;references:Id"`  // Foreign key relationship
	EventId          uint      `gorm:"not null"`                         // Foreign key field
	Event            Event     `gorm:"foreignKey:EventId;references:Id"` // Foreign key relationship
	Status           Status    `gorm:"not null;size:20"`
	EventDate        time.Time `gorm:"not null"`
	RegistrationDate time.Time `gorm:"not null"`
}

type Status string

const (
	Completed = "COMPLETED"
	Pending   = "PENDING"
	Failed    = "FAILED"
)

func (s Status) String() string {
	switch s {
	case Completed:
		return "COMPLETED"
	case Pending:
		return "PENDING"
	case Failed:
		return "FAILED"
	default:
		return "UNKNOWN"
	}
}
