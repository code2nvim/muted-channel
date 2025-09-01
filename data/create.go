package data

import (
	"log"
)

func (database *Database) CreateTables() {
	// TODO: remove all "DROP" after testing
	database.exec("DROP TABLE IF EXISTS message")
	database.exec("DROP TABLE IF EXISTS member")
	database.exec("DROP TABLE IF EXISTS room")
	database.exec("DROP TABLE IF EXISTS account")

	database.exec(`
	CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50),
		password VARCHAR(50)
	);

	CREATE TABLE IF NOT EXISTS room (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50)
	);

	CREATE TABLE IF NOT EXISTS member (
		room_id INTEGER REFERENCES room(id),
		user_id INTEGER REFERENCES account(id),
		PRIMARY KEY (room_id, user_id)
	);

	CREATE TABLE IF NOT EXISTS message (
		id SERIAL PRIMARY KEY,
		room_id INTEGER REFERENCES room(id),
		user_id INTEGER REFERENCES account(id),
		sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		content TEXT
	);
	`)
}

func (database *Database) CreateAccount(username, password string) {
	var exists bool

	database.queryRow(`
	SELECT EXISTS (
		SELECT 1 FROM account WHERE username = $1
	)
	`, username).Scan(&exists)

	if exists {
		log.Printf("error: username %s exists", username)
		return
	}

	database.exec(`
	INSERT INTO account (
		username, password
	) VALUES ($1, $2);
	`, username, password)
}

func (database *Database) CreateRoom(room string) {
	var exists bool

	database.queryRow(`
	SELECT EXISTS (
		SELECT 1 FROM room WHERE name = $1
	)
	`, room).Scan(&exists)

	if exists {
		log.Printf("error: room %s exists", room)
		return
	}

	database.exec(`
	INSERT INTO room (
		name
	) VALUES ($1);
	`, room)
}

func (database *Database) JoinRoom(user, room string) {
	var exists bool
	user_id, room_id := database.user_id(user), database.room_id(room)

	database.queryRow(`
	SELECT EXISTS (
		SELECT 1 FROM member WHERE user_id = $1 AND room_id = $2
	)
	`, user_id, room_id).Scan(&exists)

	if exists {
		log.Printf("error: %s is in %s already", user, room)
		return
	}

	database.exec(`
	INSERT INTO member (
		user_id, room_id
	) VALUES ($1, $2);
	`, user_id, room_id)
}
