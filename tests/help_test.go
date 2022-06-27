package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Food-to-Share/server/handlers"
	"github.com/Food-to-Share/server/models"
	"github.com/stretchr/testify/assert"
)

func TestAddHelpHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/addHelp", handlers.AddHelp)
	help := models.Help{
		Lunch:     "TestOrg",
		Breakfast: "TestOrg",
		Dinner:    "TestOrg",
	}

	jsonValue, _ := json.Marshal(help)

	req, _ := http.NewRequest("Post", "/addHelp", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetHelpHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/getAllHelp", handlers.GetAllHelps)

	req, _ := http.NewRequest("Get", "/getAllHelp", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var helps []models.Help
	json.Unmarshal(w.Body.Bytes(), &helps)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, helps)
}

func TestDeleteHelpHandler(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/deleteHelp/:id", handlers.DelHelp)

	r.PUT("/deleteUser/:id", handlers.DelUser)

	req, _ := http.NewRequest("Delete", "/deleteHelp/"+"4c8ba029-2421-4e8b-b4c5-82ddeb866fc1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
