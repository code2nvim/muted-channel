package server

import (
	"log"
	"net/http"

	"github.com/code2nvim/muted-channel/data"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func postAccount(c *gin.Context, database *data.Database) {
	var account data.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		return
	}
	database.CreateAccount(account.Username, account.Password)
	c.JSON(http.StatusOK, gin.H{"status": "Created account!"})
}

func postLogin(c *gin.Context, database *data.Database) {
	var account data.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		return
	}

	for _, data := range database.QueryAccounts() {
		err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(data.Password))
		if account.Username == data.Username && err != nil {
			c.SetCookie("username", account.Username, 7200, "/", "", false, true)
			c.JSON(http.StatusOK, gin.H{"status": "Login successful!"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"status": "Login failed!"})
}

func postMessage(c *gin.Context, database *data.Database) {
	var message data.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		return
	}
	database.CreateMessage(message.User, message.Room, message.Content)
	c.JSON(http.StatusOK, gin.H{"status": "Send successful!"})
}
