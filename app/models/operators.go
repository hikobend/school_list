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
	Schools   []School
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

func (o *Operator) CreateSession() (session Session, err error) {
	session = Session{}
	// sessionを作成するコマンド
	cmd1 := `insert into sessions (uuid, email, operator_id, created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(
		cmd1,
		createUUID(),
		o.Email,
		o.ID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	// sessionを取得すためのコマンド
	cmd2 := `select id, uuid, email, operator_id, created_at from sessions where operator_id = ? and email = ?`

	err = Db.QueryRow(cmd2, o.ID, o.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.OperatorID,
		&session.CreatedAt)

	return session, err
}

// sessionがデータベースに存在するか確認
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, operator_id, created_at from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.OperatorID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}

	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (sess *Session) GetOperatorBySession() (operator Operator, err error) {
	operator = Operator{}
	cmd := `select id, uuid, name, email, created_at from operators where id = ?`
	err = Db.QueryRow(cmd, sess.OperatorID).Scan(
		&operator.ID,
		&operator.UUID,
		&operator.Name,
		&operator.Email,
		&operator.CreatedAt)

	return operator, err
}
