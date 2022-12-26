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
	CallID           string `json:"callID"`
	E2EE             bool   `json:"e2ee"`
	ViewerPrice      string `json:"viewerPrice"`
	ParticipantPrice string `json:"participantPrice"`
}

// Room model
type Room struct {
	gorm.Model
	SID              string
	CallID           string
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
	SID     string `json:"sid"`
	UID     string `json:"uid"`
	Name    string `json:"name"`
	IsHost  bool   `json:"isHost"`
	Account string `json:"account"`
	URL     string `json:"url"`
	CallID  string `json:"callID"`
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

func createRoom(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

		var roomRequest RoomRequest
		err := c.Bind(&roomRequest)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		log.Printf("createRoom: %v", roomRequest)

		SID := shortuuid.New()
		UID := shortuuid.New()
		key := generateKey()

		url, err := getNodeURL()
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		token := &Token{
			SID:     SID,
			UID:     UID,
			Name:    roomRequest.Name,
			IsHost:  true,
			Account: os.Getenv("NEAR_ACCOUNT"),
			URL:     os.Getenv("CALLBACK_URL"),
			CallID:  roomRequest.CallID,
		}

		tokenString, signature, err := getTokenSignature(token)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		room := &Room{
			SID:              SID,
			CallID:           roomRequest.CallID,
			Key:              key,
			Title:            roomRequest.Title,
			HostUID:          UID,
			E2EE:             roomRequest.E2EE,
			ViewerPrice:      roomRequest.ViewerPrice,
			ParticipantPrice: roomRequest.ParticipantPrice,
		}
		db.Create(&room)

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

		UID := shortuuid.New()
		url, err := getNodeURL()
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		token := &Token{
			SID:     roomRequest.SID,
			UID:     UID,
			Name:    roomRequest.Name,
			IsHost:  false,
			Account: os.Getenv("NEAR_ACCOUNT"),
			URL:     os.Getenv("CALLBACK_URL"),
			CallID:  room.CallID,
		}

		tokenString, signature, err := getTokenSignature(token)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
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
		return c.JSON(http.StatusOK, "")
	}
}
