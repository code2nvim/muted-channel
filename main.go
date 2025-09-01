package main

import (
	"github.com/code2nvim/muted-channel/data"
	"github.com/code2nvim/muted-channel/server"
)

func main() {
	data := data.Database{
		DB: data.Conn(".env"),
	}
	defer data.DB.Close()
	data.CreateTables()
	data.CreateAccount("Foo", "Bar")
	data.CreateRoom("channel1")
	data.CreateRoom("channel2")
	data.JoinRoom("Foo", "channel1")
	data.JoinRoom("Foo", "channel2")
	server.Route(&data)
}
