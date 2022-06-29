package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

type NewUser struct {
	Jid         string `json:"jid"`
	DisplayName string `json:"displayName"`
}

type OtherUserInfo struct {
	MXID string `json:"mxid"`
	JID  string `json:"jid"`
	Name string `json:"displayname"`
}

type BridgeStartUserResp struct {
	RoomID      string        `json:"room_id"`
	OtherUser   OtherUserInfo `json:"other_user"`
	JustCreated bool          `json:"just_created"`
}

type SyncNissReq struct {
	Jid  string `json:"jid"`
	Niss string `json:"niss"`
}

func AddUser(c *gin.Context) {

	db := database.GetDatabase()

	var BridgeServerURL = "http://localhost:29318"

	httpposturl := BridgeServerURL + "/_matrix/provision/v1/startUser"

	var user models.User
	var help models.Help

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	errHelp := c.ShouldBindJSON(&help)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + errHelp.Error(),
		})
		return
	}

	user.ID = uuid.NewString()
	help.ID = user.ID

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}

	errHelp = db.Create(&help).Error

	if errHelp != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create help: " + errHelp.Error(),
		})
		return
	}

	newUsers := NewUser{
		Jid:         user.ID,
		DisplayName: user.Name,
	}

	data, err := json.Marshal(newUsers)
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

		var br BridgeStartUserResp

		err = json.Unmarshal(respbody, &br)

		httpposturl = BridgeServerURL + "/_matrix/provision/v1/syncNiss"

		syncNiss := SyncNissReq{
			Jid:  br.OtherUser.JID,
			Niss: user.Ssn,
		}

		data, err := json.Marshal(syncNiss)
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
			c.JSON(http.StatusCreated, br)
		}
	}
	//defer response.Body.Close()
}
