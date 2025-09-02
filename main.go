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
	data.CreateRoom("channel1")
	data.CreateRoom("channel2")
	data.JoinRoom("Foo", "channel1")
	data.JoinRoom("Foo", "channel2")
	data.SendMessage("Foo", "channel2", "hello")
	data.SendMessage("Foo", "channel2", "world")
	server.Route(&data)
}
