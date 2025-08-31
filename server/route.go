package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		cookie, err := c.Cookie("login_status")
		if err != nil {
			cookie = "failed"
			c.SetCookie("login_status", "success", 7200, "/", "", true, true)
		}
		log.Println("login_status: ", cookie)
	})
	router.RunTLS(":8088", ".local/cert.pem", ".local/key.pem")
}
