package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func VerifyUserNiss(c *gin.Context) {

	var user models.User

	var BridgeServerURL = "http://localhost:29318"

	data, err := json.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}

	httpposturl := BridgeServerURL
	reader := bytes.NewBuffer(data)

	response, errorPost := http.Post(httpposturl, "application/json", reader)
	if errorPost != nil {
		panic(errorPost)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		respbody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		c.JSON(200, respbody)
	}
}
