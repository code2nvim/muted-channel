package data

func (database *Database) user(id int) string {
	var name string
	err := database.queryRow("SELECT username FROM account WHERE id = $1", id).Scan(&name)
	check(err)
	return name
}

func (database *Database) user_id(user string) int {
	var id int
	err := database.queryRow("SELECT id FROM account WHERE username = $1", user).Scan(&id)
	check(err)
	return id
}

func (database *Database) room(id int) string {
	var name string
	err := database.queryRow("SELECT name FROM room WHERE id = $1", id).Scan(&name)
	check(err)
	return name
}

func (database *Database) room_id(room string) int {
	var id int
	err := database.queryRow("SELECT id FROM room WHERE name = $1", room).Scan(&id)
	check(err)
	return id
}

func (database *Database) QueryRooms() []Room {
	rows := database.query("SELECT * FROM room")
	defer rows.Close()

	var rooms []Room
	for rows.Next() {
		var room Room
		err := rows.Scan(&room.ID, &room.Name)
		check(err)
		rooms = append(rooms, room)
	}
	check(rows.Err())

	return rooms
}

func (database *Database) QueryMessages(room string) []Message {
	rows := database.query("SELECT * FROM message WHERE room_id = $1", database.room_id(room))
	defer rows.Close()

	var messages []Message
	var user_id, room_id int
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.ID, &user_id, &room_id, &message.SentAt, &message.Content)
		message.User, message.Room = database.user(user_id), database.room(room_id)
		check(err)
		messages = append(messages, message)
	}
	check(rows.Err())
	return messages
}
