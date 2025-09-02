package server

import (
	"encoding/json"
	"net/http"

	"github.com/code2nvim/muted-channel/data"
	"github.com/gin-gonic/gin"
)

func getUsername(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"username": ""})
	}
	c.JSON(http.StatusOK, gin.H{"username": username})
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
