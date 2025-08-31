package data

import (
	"log"
)

func (data *Data) CreateTables() {
	// data.exec("DROP TABLE IF EXISTS account;")
	data.exec(`
	CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50),
		password VARCHAR(50)
	);`)
}

func (data *Data) CreateAccount(username, password string) {
	exists := false
	data.queryRow(`
	SELECT EXISTS (
		SELECT 1
		FROM account
		WHERE username = $1
	)`, username).Scan(&exists)
	if exists {
		log.Printf("error: %s exists", username)
		return
	}
	data.exec(`
	INSERT INTO account (
		username, password
	) VALUES ($1, $2);`, username, password)
}
