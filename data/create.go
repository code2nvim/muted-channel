package data

import (
	"log"
)

func (data *Data) CreateTables() {
	// TODO: remove all "DROP" after testing
	data.exec("DROP TABLE IF EXISTS message")
	data.exec("DROP TABLE IF EXISTS member")
	data.exec("DROP TABLE IF EXISTS room")
	data.exec("DROP TABLE IF EXISTS account")

	data.exec(`
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
		sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
}

func (data *Data) CreateAccount(username, password string) {
	var exists bool

	data.queryRow(`
	SELECT EXISTS (
		SELECT 1 FROM account WHERE username = $1
	)
	`, username).Scan(&exists)

	if exists {
		log.Printf("error: username %s exists", username)
		return
	}

	data.exec(`
	INSERT INTO account (
		username, password
	) VALUES ($1, $2);
	`, username, password)
}

func (data *Data) CreateRoom(room string) {
	var exists bool

	data.queryRow(`
	SELECT EXISTS (
		SELECT 1 FROM room WHERE name = $1
	)
	`, room).Scan(&exists)

	if exists {
		log.Printf("error: room %s exists", room)
		return
	}

	data.exec(`
	INSERT INTO room (
		name
	) VALUES ($1);
	`, room)
}

func (data *Data) JoinRoom(user, room string) {
	var exists bool
	user_id, room_id := data.user_id(user), data.room_id(room)

	data.queryRow(`
	SELECT EXISTS (
		SELECT 1 FROM member WHERE user_id = $1 AND room_id = $2
	)
	`, user_id, room_id).Scan(&exists)

	if exists {
		log.Printf("error: %s is in %s already", user, room)
		return
	}

	data.exec(`
	INSERT INTO member (
		user_id, room_id
	) VALUES ($1, $2);
	`, user_id, room_id)
}
