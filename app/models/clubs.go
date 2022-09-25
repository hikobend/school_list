package models

import (
	"log"
	"time"
)

type Club struct {
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

func GetClub(id int) (club Club, err error) {
	cmd := `select 
					id,
					name,
					content,
					school_id,
					created_at from clubs where id = ?`

	club = Club{}

	err = Db.QueryRow(cmd, id).Scan(
		&club.ID,
		&club.Name,
		&club.Content,
		&club.SchoolID,
		&club.CreatedAt,
	)

	if err != nil {
		log.Fatalln(err)
	}

	return club, err
}
