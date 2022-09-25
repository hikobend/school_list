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

// 1つのclassを取得
func GetClass(id int) (class Class, err error) {
	cmd := `select 
					id, 
					class_number, 
					school_id, 
					created_at from classes where id = ?`

	class = Class{}

	err = Db.QueryRow(cmd, id).Scan(
		&class.ID,
		&class.ClassNumber,
		&class.SchoolID,
		&class.CreatedAt,
	)

	if err != nil {
		log.Fatalln(err)
	}

	return class, err
}

// 複数のclassを取得
func GetClasses() (classes []Class, err error) {
	cmd := `select 
					id, 
					class_number, 
					school_id, 
					created_at from classes`
	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var class Class

		err = rows.Scan(
			&class.ID,
			&class.ClassNumber,
			&class.SchoolID,
			&class.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		classes = append(classes, class)
	}
	rows.Close()

	return classes, err
}
