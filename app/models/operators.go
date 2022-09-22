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

func GetOperator(id int) (operator Operator, err error) {
	operator = Operator{}
	cmd := `select id, uuid, name, email, password, created_at from operators where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&operator.ID,
		&operator.UUID,
		&operator.Name,
		&operator.Email,
		&operator.PasaWord,
		&operator.CreatedAt)

	return operator, err
}
