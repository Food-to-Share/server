package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Food-to-Share/server/handlers"
	"github.com/Food-to-Share/server/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/getUsers", handlers.GetAllUsers)
	req, _ := http.NewRequest("GET", "/getUsers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
}

func TestGetUserByIdHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/getUsersById/:id", handlers.GetUserById)
	req, _ := http.NewRequest("GET", "/getUsersById/"+"4c8ba029-2421-4e8b-b4c5-82ddeb866fc1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	mockUser := models.User{
		ID:           "4c8ba029-2421-4e8b-b4c5-82ddeb866fc1",
		Ssn:          "14698712379",
		Username:     "pedrot",
		Name:         "Pedro tinoco",
		Email:        "pedrot@gmail.com",
		Contact:      937765982,
		Address:      "Rua caridade 2",
		Organization: "Help4You",
	}

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockUser, user)
}

func TestAddUserHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/addUser", handlers.AddUser)
	user := models.User{
		Username:     "teste",
		Name:         "João Saraiva",
		Email:        "joaotest@gmail.com",
		Organization: "FoodToShare",
		Contact:      911345123,
		Address:      "Rua Caridade 3",
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("Post", "/addUser", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateUserHandler(t *testing.T) {
	r := SetUpRouter()
	r.PUT("/updateUser", handlers.UpdateUser)
	user := models.User{
		ID:           "df471b07-31a6-4304-921a-ba0cacf36c8d",
		Name:         "João Maria",
		Username:     "kevind",
		Organization: "FoodToShare",
		Email:        "kevintest@gmail.com",
		Contact:      911345123,
		Address:      "Rua Caridade 3",
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("Put", "/updateUser", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(responseData), string(jsonValue), "error message %s", "formatted")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUserHandler(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/deleteUser/:id", handlers.DelUser)

	req, _ := http.NewRequest("Delete", "/deleteUser/"+"4c8ba029-2421-4e8b-b4c5-82ddeb866fc1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestVerifyNissHandler(t *testing.T) {
	mockResponse := `{""}`
	r := SetUpRouter()
	r.POST("/verifyUser", handlers.VerifyUserNiss)
	niss := models.User{
		Ssn: "13547892186",
	}

	jsonValue, _ := json.Marshal(niss)

	req, _ := http.NewRequest("POST", "/verifyUser", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Containsf(t, string(responseData), mockResponse, "error message %s", "formatted")
	assert.Equal(t, http.StatusOK, w.Code)
}
