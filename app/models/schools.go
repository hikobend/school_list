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
		o.Name,
		o.ID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
