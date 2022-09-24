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
	PassWord  string
	CreatedAt time.Time
}

type Session struct {
	ID         int
	UUID       string
	Email      string
	OperatorID int
	CreatedAt  time.Time
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
		Encrypt(o.PassWord),
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
		&operator.PassWord,
		&operator.CreatedAt)

	return operator, err
}

func (o *Operator) UpdateOperator() (err error) {
	cmd := `update operators set name = ?, email = ? where id = ?`

	_, err = Db.Exec(cmd, o.Name, o.Email, o.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (o *Operator) DeleteOperator() (err error) {
	cmd := `delete from operators where id = ?`

	_, err = Db.Exec(cmd, o.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetOperatorByEmail(email string) (operator Operator, err error) {
	operator = Operator{}
	cmd := `select id, uuid, name, email, password, created_at from operators where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&operator.ID,
		&operator.UUID,
		&operator.Name,
		&operator.Email,
		&operator.PassWord,
		&operator.CreatedAt)

	return operator, err
}
