package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://localhost:5173"}
	config.AllowCredentials = true
	return cors.New(config)
}
