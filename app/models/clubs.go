package models

import "time"

type Clob struct {
	ID        int
	Name      string
	Content   string
	SchoolID  int
	CreatedAt time.Time
}
