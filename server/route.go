package server

import (
	"net/http"

	"github.com/code2nvim/muted-channel/data"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, database *data.Database) {
	router.Use(cors.Default())

	router.GET("/api/rooms", func(c *gin.Context) { getRooms(c, database) })
	router.GET("/api/room/:name", func(c *gin.Context) { getRoom(c, database) })
	router.GET("/api/username", func(c *gin.Context) { getUsername(c) })

	router.POST("/api/account", func(c *gin.Context) { postAccount(c, database) })
	router.POST("/api/login", func(c *gin.Context) { postLogin(c, database) })
	router.POST("/api/message", func(c *gin.Context) { postMessage(c, database) })

	router.StaticFS("/assets", http.Dir("./dist/assets"))
	router.NoRoute(func(c *gin.Context) { c.File("./dist/index.html") })

	router.RunTLS(":8088", ".local/cert.pem", ".local/key.pem")
}
