package models

import (
	"log"
	"time"
)

type Class struct {
	ID          int
	ClassNumber string
	SchoolID    int
	CreatedAt   time.Time
}

func (s *School) CreateClass(class_number string) (err error) {
	cmd := `insert into classes(
		class_number, 
		school_id, 
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd,
		class_number,
		s.ID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
