package models

import (
	"log"
	"time"
)

type Operator struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PasaWord  string
	CreatedAt time.Time
}

func (o *Operator) CreateOperator() (err error) {
	cmd := `insert into operators(
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		o.Name,
		o.Email,
		Encrypt(o.PasaWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
