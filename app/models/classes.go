package models

import "time"

type Class struct {
	ID          int
	ClassNumber string
	SchoolID    int
	CreatedAt   time.Time
}
