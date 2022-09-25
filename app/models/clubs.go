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

func GetClubs() (clubs []Club, err error) {
	cmd := `select 
					id,
					name,
					content,
					school_id,
					created_at from clubs`

	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var club Club

		err = rows.Scan(
			&club.ID,
			&club.Name,
			&club.Content,
			&club.SchoolID,
			&club.CreatedAt,
		)

		if err != nil {
			log.Fatalln(err)
		}

		clubs = append(clubs, club)
	}
	rows.Close()

	return clubs, err
}

func (s *School) GetClubBySchool() (clubs []Club, err error) {
	cmd := `select
					id,
					name,
					content,
					school_id,
					created_at from clubs where school_id = ?`

	rows, err := Db.Query(cmd, s.ID)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var club Club

		err = rows.Scan(
			&club.ID,
			&club.Name,
			&club.Content,
			&club.SchoolID,
			&club.CreatedAt,
		)

		if err != nil {
			log.Fatalln(err)
		}

		clubs = append(clubs, club)
	}
	rows.Close()

	return clubs, err
}

func (c *Club) UpdateClub() (err error) {
	cmd := `update clubs set
					name = ?,
					content = ?,
					school_id = ? where id = ?`

	_, err = Db.Exec(
		cmd,
		c.Name,
		c.Content,
		c.SchoolID,
		c.ID,
	)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
