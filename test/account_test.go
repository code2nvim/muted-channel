package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/code2nvim/muted-channel/data"
	"github.com/code2nvim/muted-channel/server"
	"github.com/gin-gonic/gin"
)

type account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func setup() (*data.Database, *gin.Engine) {
	data := data.Database{
		DB: data.Conn(
			fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_SSLMODE"),
			)),
	}
	data.CreateTables()

	router := gin.Default()
	server.Route(router, &data)

	return &data, router
}

func TestAccount(t *testing.T) {
	data, router := setup()
	defer data.DB.Close()

	t.Run("Create Account", func(t *testing.T) { runCreateTest(t, router) })
	t.Run("Login Account", func(t *testing.T) { runLogin(t, router) })
	t.Run("Login Account Failed", func(t *testing.T) { runLoginFailed(t, router) })
}

func runCreateTest(t *testing.T, router *gin.Engine) {
	created, _ := json.Marshal(account{
		Username: "Test",
		Password: "Real",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/account", strings.NewReader(string(created)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status: %d", w.Code)
	}
}

func runLogin(t *testing.T, router *gin.Engine) {
	created, _ := json.Marshal(account{
		Username: "Test",
		Password: "Real",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(string(created)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status: %d", w.Code)
	}
}

func runLoginFailed(t *testing.T, router *gin.Engine) {
	created, _ := json.Marshal(account{
		Username: "Test",
		Password: "Fake",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(string(created)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Status: %d", w.Code)
	}
}
