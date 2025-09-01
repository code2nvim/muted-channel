package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/code2nvim/muted-channel/data"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route(database *data.Database) {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/login", func(c *gin.Context) { getLogin(c) })
	router.GET("/rooms", func(c *gin.Context) { getRooms(c, database) })
	router.GET("/room/:name", func(c *gin.Context) { getRoom(c, database) })
	router.RunTLS(":8088", ".local/cert.pem", ".local/key.pem")
}

func getLogin(c *gin.Context) {
	cookie, err := c.Cookie("login_status")
	if err != nil {
		cookie = "failed"
		c.SetCookie("login_status", "success", 7200, "/", "", true, true)
	}
	log.Println("login_status: ", cookie)

}

func getRooms(c *gin.Context, database *data.Database) {
	first := database.QueryRooms()
	get := func(rooms []data.Room) {
		json, _ := json.Marshal(rooms)
		c.SSEvent("rooms", json)
		c.Writer.(http.Flusher).Flush()
	}
	get(first)
	for {
		rooms := database.QueryRooms()
		if len(rooms) != len(first) {
			get(rooms)
		}
	}
}

func getRoom(c *gin.Context, database *data.Database) {
	room := c.Param("name")
	first := database.QueryMessages(room)
	get := func(messages []data.Message) {
		json, _ := json.Marshal(messages)
		c.SSEvent("messages", json)
		c.Writer.(http.Flusher).Flush()
	}
	get(first)
	for {
		messages := database.QueryMessages(room)
		if len(messages) != len(first) {
			get(messages)
		}
	}
}
