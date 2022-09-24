package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"example.com/school/config"
	// ドライバのインストール
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

// エラーの宣言
var err error

// テーブル名の宣言
const (
	tableNameOperator = "operators"
	tableNameSchool   = "schools"
	tableNameSession  = "sessionss"
)

// テーブルはmain関数の前に作成
func init() {
	// データーベースとエラー。ドライバとデーターベース名
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}

	cmdO := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameOperator)

	Db.Exec(cmdO)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING,
		operator_id STRING,
		created_at DATETIME)`, tableNameSchool)

	Db.Exec(cmdS)
}

// UUID作成
// 返り値をuuidobj
func createUUID() (uuidobj uuid.UUID) {
	// NewUUIDを使用
	uuidobj, _ = uuid.NewUUID()
	// returnで返す
	return uuidobj
}

// パスワード作成
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	// returnで返す
	return cryptext
}
