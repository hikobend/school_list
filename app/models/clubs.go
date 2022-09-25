package models

import (
	"log"
	"time"
)

type Clob struct {
	ID        int
	Name      string
	Content   string
	SchoolID  int
	CreatedAt time.Time
}

func (s *School) CreateClub(name string, content string) (err error) {
	cmd := `insert into clubs(
		name,
		content,
		school_id,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		name,
		content,
		s.ID,
		time.Now(),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
