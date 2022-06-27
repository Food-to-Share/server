package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/handlers"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	database.StartDB()
	router := gin.Default()

	return router
}

func TestLoginHandler(t *testing.T) {
	mockResponse := `"token"`
	r := SetUpRouter()
	r.POST("/login", handlers.Login)
	login := models.Login{
		Email:    "kevind@ipvc.pt",
		Password: "1234",
	}

	jsonValue, _ := json.Marshal(login)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Containsf(t, string(responseData), mockResponse, "error message %s", "formatted")
	assert.Equal(t, http.StatusOK, w.Code)
}
