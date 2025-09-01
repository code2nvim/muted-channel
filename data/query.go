package data

func (database *Database) user_id(user string) int {
	var id int
	err := database.queryRow(`SELECT id FROM account WHERE username = $1`, user).Scan(&id)
	check(err)
	return id
}

func (database *Database) room_id(room string) int {
	var id int
	err := database.queryRow(`SELECT id FROM room WHERE name = $1`, room).Scan(&id)
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
