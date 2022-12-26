package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

// Participant model
type Participant struct {
	gorm.Model
	Name   string
	UID    string
	SID    string
	IsHost bool
}

// RoomRequest model
type RoomRequest struct {
	Name             string `json:"name"`
	UID              string `json:"uid"`
	SID              string `json:"sid"`
	Title            string `json:"title"`
	E2EE             bool   `json:"e2ee"`
	ViewerPrice      string `json:"viewerPrice"`
	ParticipantPrice string `json:"participantPrice"`
	NoPublish        bool   `json:"noPublish"`
}

// Room model
type Room struct {
	gorm.Model
	SID              string
	Key              string
	Title            string
	HostUID          string
	E2EE             bool
	ViewerPrice      string
	ParticipantPrice string
}

// RoomView model
type RoomView struct {
	Title            string `json:"title"`
	HostName         string `json:"hostName"`
	Count            int    `json:"count"`
	E2EE             bool   `json:"e2ee"`
	ViewerPrice      string `json:"viewerPrice"`
	ParticipantPrice string `json:"participantPrice"`
}

// Token model
type Token struct {
	SID       string `json:"sid"`
	UID       string `json:"uid"`
	Name      string `json:"name"`
	IsHost    bool   `json:"isHost"`
	Account   string `json:"account"`
	URL       string `json:"url"`
	CallID    string `json:"callID"`
	NoPublish bool   `json:"noPublish"`
}

// TokenView model
type TokenView struct {
	Token     string `json:"token"`
	Signature string `json:"signature"`
	URL       string `json:"url"`
	SID       string `json:"sid"`
	UID       string `json:"uid"`
	Key       string `json:"key"`
}

//NotifyData data
type NotifyData struct {
	Duration int    `json:"duration"`
	SID      string `json:"sid"`
	UID      string `json:"uid"`
	Type     string `json:"type"`
}

//NotifyRequest data
type NotifyRequest struct {
	Data      []byte `json:"data"`
	Signature []byte `json:"signature"`
}

//NotifyResponse data
type NotifyResponse struct {
	Epoch     uint64 `json:"epoch"`
	Signature []byte `json:"signature"`
}

// Call model
type Call struct {
	gorm.Model
	SID    string
	CallID string
	NodeID string
}

func createRoom(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

		var roomRequest RoomRequest
		err := c.Bind(&roomRequest)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		log.Printf("createRoom: %v", roomRequest)

		SID := shortuuid.New()
		UID := shortuuid.New()
		CallID := shortuuid.New()
		key := generateKey()

		url, nodeID, err := getNodeURL()
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		token := &Token{
			SID:       SID,
			UID:       UID,
			Name:      roomRequest.Name,
			IsHost:    true,
			Account:   os.Getenv("NEAR_ACCOUNT"),
			URL:       os.Getenv("CALLBACK_URL"),
			CallID:    CallID,
			NoPublish: false,
		}

		tokenString, signature, err := getTokenSignature(token)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		room := &Room{
			SID:              SID,
			Key:              key,
			Title:            roomRequest.Title,
			HostUID:          UID,
			E2EE:             roomRequest.E2EE,
			ViewerPrice:      roomRequest.ViewerPrice,
			ParticipantPrice: roomRequest.ParticipantPrice,
		}
		db.Create(&room)

		call := &Call{
			SID:    SID,
			CallID: CallID,
			NodeID: nodeID,
		}
		db.Create(&call)

		participant := &Participant{
			Name:   roomRequest.Name,
			SID:    SID,
			UID:    UID,
			IsHost: true,
		}
		db.Create(&participant)

		tokenView := &TokenView{
			Token:     tokenString,
			Signature: signature,
			URL:       url,
			SID:       SID,
			UID:       UID,
			Key:       key,
		}

		return c.JSON(http.StatusOK, tokenView)
	}
}

func joinRoom(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

		var roomRequest RoomRequest
		err := c.Bind(&roomRequest)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		log.Printf("joinRoom: %v", roomRequest)

		if roomRequest.SID == "" {
			return c.String(http.StatusBadRequest, "SID required")
		}

		var room Room
		db.Where("s_id=?", roomRequest.SID).First(&room)
		if room.SID != roomRequest.SID {
			return c.String(http.StatusNotFound, "")
		}

		url, nodeID, err := getNodeURL()
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var call Call
		db.Where("s_id=? AND node_id=?", roomRequest.SID, nodeID).First(&call)
		if call.SID != roomRequest.SID {
			call = Call{
				SID:    roomRequest.SID,
				CallID: shortuuid.New(),
				NodeID: nodeID,
			}
		}

		UID := shortuuid.New()

		token := &Token{
			SID:       roomRequest.SID,
			UID:       UID,
			Name:      roomRequest.Name,
			IsHost:    false,
			Account:   os.Getenv("NEAR_ACCOUNT"),
			URL:       os.Getenv("CALLBACK_URL"),
			CallID:    call.CallID,
			NoPublish: roomRequest.NoPublish,
		}

		tokenString, signature, err := getTokenSignature(token)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		participant := &Participant{
			Name:   roomRequest.Name,
			SID:    roomRequest.SID,
			UID:    UID,
			IsHost: false,
		}
		db.Create(&participant)

		tokenView := &TokenView{
			Token:     tokenString,
			Signature: signature,
			URL:       url,
			SID:       roomRequest.SID,
			UID:       UID,
			Key:       room.Key,
		}

		return c.JSON(http.StatusOK, tokenView)
	}
}

func infoRoom(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var roomRequest RoomRequest
		err := c.Bind(&roomRequest)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		log.Printf("joinRoom: %v", roomRequest)

		if roomRequest.SID == "" {
			return c.String(http.StatusBadRequest, "SID required")
		}

		var room Room
		db.Where("s_id=?", roomRequest.SID).First(&room)
		if room.SID != roomRequest.SID {
			return c.String(http.StatusNotFound, "")
		}

		var participants []Participant
		db.Where("s_id=?", room.SID).Find(&participants)

		var host Participant
		db.Where("uid=?", room.HostUID).First(&host)

		roomView := &RoomView{
			Title:            room.Title,
			Count:            len(participants),
			HostName:         host.Name,
			E2EE:             room.E2EE,
			ViewerPrice:      room.ViewerPrice,
			ParticipantPrice: room.ParticipantPrice,
		}

		return c.JSON(http.StatusOK, roomView)
	}
}

func callbackRoom(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var notifyRequest NotifyRequest
		err := c.Bind(&notifyRequest)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		log.Printf("callbackRoom: %v", notifyRequest)

		// var room Room
		// db.Where("s_id=?", notifyData.SID).First(&room)
		// if room.SID != notifyData.SID {
		// 	return c.String(http.StatusNotFound, "")
		// }

		// if notifyData.Type == "join" {
		// 	var participant Participant
		// 	db.Where("uid=?", notifyData.UID).First(&participant)
		// 	if participant.UID != notifyData.UID {
		// 		return c.String(http.StatusNotFound, "")
		// 	}
		// }
		return c.String(http.StatusNotFound, "")
	}
}
