package main

import (
	"fmt"
	"os"

	"github.com/code2nvim/muted-channel/data"
	"github.com/code2nvim/muted-channel/server"
)

func main() {
	data := data.Database{
		DB: data.Conn(
			fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_SSLMODE"),
			)),
	}
	defer data.DB.Close()

	data.CreateTables()
	data.CreateAccount("Foo", "Bar")
	data.CreateRoom("default channel")
	data.CreateRoom("channel 1")
	data.CreateRoom("channel 2")
	data.CreateRoom("example")
	data.JoinRoom("Foo", "channel 1")
	data.JoinRoom("Foo", "example")
	data.SendMessage("Foo", "default channel", "default")
	data.SendMessage("Foo", "example", "Hello")
	data.SendMessage("Foo", "example", "World")

	server.Route(&data)
}
