package data

import (
	"database/sql"
	"os"
	"time"

	"github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID      int       `json:"id"`
	User    string    `json:"user"`
	Room    string    `json:"room"`
	SentAt  time.Time `json:"sent_at"`
	Content string    `json:"content"`
}

func Conn(env string) *sql.DB {
	text, err := os.ReadFile(env)
	check(err)
	connector, err := pq.NewConnector(string(text))
	check(err)
	db := sql.OpenDB(connector)
	check(err)
	return db
}

func (database *Database) exec(query string, args ...any) {
	_, err := database.DB.Exec(query, args...)
	check(err)
}

func (database *Database) query(query string, args ...any) *sql.Rows {
	rows, err := database.DB.Query(query, args...)
	check(err)
	return rows
}

func (database *Database) queryRow(query string, args ...any) *sql.Row {
	row := database.DB.QueryRow(query, args...)
	return row
}
