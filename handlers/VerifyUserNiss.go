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

type NissReq struct {
	Niss string `json:"niss"`
}

type LookupResp struct {
	Medium  string `json:"medium"`
	Address string `json:"address"`
	MXID    string `json:"mxid"`
}

func VerifyUserNiss(c *gin.Context) {

	var user models.User

	var BridgeServerURL = "http://localhost:29318"

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}

	httpposturl := BridgeServerURL + "/_matrix/provision/v1/lookup"

	nissreq := NissReq{
		Niss: user.Ssn,
	}

	data, err := json.Marshal(nissreq)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)

	request, errorPost := http.NewRequest("POST", httpposturl, reader)
	if errorPost != nil {
		panic(errorPost)
	}
	client := &http.Client{}

	response, error := client.Do(request)
	if error != nil {
		c.JSON(400, gin.H{
			"error": "Cannot do request: " + error.Error(),
		})
		return
	}

	if response.StatusCode == http.StatusOK {
		respbody, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		var lr LookupResp

		err = json.Unmarshal(respbody, &lr)
		c.JSON(response.StatusCode, lr)
	}
	// defer response.Body.Close()

}
