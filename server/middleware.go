package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return cors.New(config)
}
