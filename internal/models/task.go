package models

import (
	"time"
)

type Task struct {
	Id          string
	Username    string
	Title       string
	Description string
	Status      string
	StartDate   time.Time
	EndDate     time.Time
	CreatedDate time.Time
}

func Print() string {
	return "shadow"
}
