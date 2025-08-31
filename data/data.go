package data

import (
	"database/sql"
	"os"

	"github.com/lib/pq"
)

type Data struct {
	DB *sql.DB
}

func connEnv() string {
	text, err := os.ReadFile(".env")
	check(err)
	return string(text)
}

func Conn() *sql.DB {
	connector, err := pq.NewConnector(connEnv())
	check(err)
	db := sql.OpenDB(connector)
	check(err)
	return db
}

func (data *Data) exec(query string, args ...any) {
	_, err := data.DB.Exec(query, args...)
	check(err)
}

func (data *Data) query(query string, args ...any) *sql.Rows {
	rows, err := data.DB.Query(query, args...)
	check(err)
	return rows
}

func (data *Data) queryRow(query string, args ...any) *sql.Row {
	row := data.DB.QueryRow(query, args...)
	return row
}
