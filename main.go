package main

import (
	"github.com/code2nvim/muted-channel/data"
)

func main() {
	data := data.Data{
		DB: data.Conn(),
	}
	defer data.DB.Close()
	data.CreateTables()
	data.CreateAccount("Foo", "Bar")
}
