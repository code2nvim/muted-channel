package data

func (data *Data) user_id(user string) int {
	var id int
	err := data.queryRow(`SELECT id FROM account WHERE username = $1`, user).Scan(&id)
	check(err)
	return id
}

func (data *Data) room_id(room string) int {
	var id int
	err := data.queryRow(`SELECT id FROM room WHERE name = $1`, room).Scan(&id)
	check(err)
	return id
}
