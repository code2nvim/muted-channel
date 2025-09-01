package server

import (
	"log"

	"github.com/code2nvim/muted-channel/data"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route(data *data.Data) {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/login", func(c *gin.Context) { getLogin(c) })
	router.GET("/rooms", func(c *gin.Context) { getRooms(c, data) })
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

func getRooms(c *gin.Context, data *data.Data) {
	rooms := data.QueryRooms()
	c.JSON(200, rooms)
}
