package models

import (
	"log"
	"time"
)

type School struct {
	ID         int
	Name       string
	OperatorID int
	CreatedAt  time.Time
}

func (o *Operator) CreateSchool(name string) (err error) {
	cmd := `insert into schools(
		name, 
		operator_id, 
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd,
		name,
		o.ID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetSchool(id int) (school School, err error) {
	cmd := `select id, name, operator_id, created_at from schools where id = ?`

	school = School{}

	err = Db.QueryRow(cmd, id).Scan(
		&school.ID,
		&school.Name,
		&school.OperatorID,
		&school.CreatedAt)

	if err != nil {
		log.Fatalln(err)
	}

	return school, err
}

func GetSchools() (schools []School, err error) {
	cmd := `select id, name, operator_id, created_at from schools`
	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	// 取り出す
	for rows.Next() {
		var school School
		err = rows.Scan(
			&school.ID,
			&school.Name,
			&school.OperatorID,
			&school.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		schools = append(schools, school)
	}
	rows.Close()

	return schools, err
}

func (o *Operator) GetSchoolByOperator() (schools []School, err error) {
	cmd := `select id, name, operator_id, created_at from schools where operator_id = ?`
	rows, err := Db.Query(cmd, o.ID)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var school School
		err = rows.Scan(
			&school.ID,
			&school.Name,
			&school.OperatorID,
			&school.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		schools = append(schools, school)
	}
	rows.Close()

	return schools, err
}

func (s *School) UpdateSchool() (err error) {
	cmd := `update schools set name = ?, operator_id = ? where id = ?`

	_, err = Db.Exec(cmd,
		s.Name,
		s.OperatorID,
		s.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *School) DeleteSchool() (err error) {
	cmd := `delete from schools where id = ?`

	_, err = Db.Exec(cmd, s.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
