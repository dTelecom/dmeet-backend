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

// ParticipantView model
type ParticipantView struct {
	Name   string `json:"name"`
	UID    string `json:"uid"`
	SID    string `json:"sid"`
	Title  string `json:"title"`
	CallID string `json:"callID"`
	E2EE   bool   `json:"e2ee"`
}

// Room model
type Room struct {
	gorm.Model
	SID     string
	CallID  string
	Key     string
	Title   string
	HostUID string
	E2EE    bool
}

// RoomView model
type RoomView struct {
	Title    string `json:"title"`
	HostName string `json:"hostName"`
	Count    int    `json:"count"`
	E2EE     bool   `json:"e2ee"`
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

		var participantView ParticipantView
		err := c.Bind(&participantView)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		log.Printf("createRoom: %v", participantView)

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
			Name:    participantView.Name,
			IsHost:  true,
			Account: os.Getenv("NEAR_ACCOUNT"),
			URL:     os.Getenv("CALLBACK_URL"),
			CallID:  participantView.CallID,
		}

		tokenString, signature, err := getTokenSignature(token)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		room := &Room{
			SID:     SID,
			CallID:  participantView.CallID,
			Key:     key,
			Title:   participantView.Title,
			HostUID: UID,
			E2EE:    participantView.E2EE,
		}
		db.Create(&room)

		participant := &Participant{
			Name:   participantView.Name,
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

		var participantView ParticipantView
		err := c.Bind(&participantView)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		log.Printf("joinRoom: %v", participantView)

		if participantView.SID == "" {
			return c.String(http.StatusBadRequest, "SID required")
		}

		var room Room
		db.Where("s_id=?", participantView.SID).First(&room)
		if room.SID != participantView.SID {
			return c.String(http.StatusNotFound, "")
		}

		UID := shortuuid.New()
		url, err := getNodeURL()
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		token := &Token{
			SID:     participantView.SID,
			UID:     UID,
			Name:    participantView.Name,
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
			Name:   participantView.Name,
			SID:    participantView.SID,
			UID:    UID,
			IsHost: false,
		}
		db.Create(&participant)

		tokenView := &TokenView{
			Token:     tokenString,
			Signature: signature,
			URL:       url,
			SID:       participantView.SID,
			UID:       UID,
			Key:       room.Key,
		}

		return c.JSON(http.StatusOK, tokenView)
	}
}

func infoRoom(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var participantView ParticipantView
		err := c.Bind(&participantView)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		log.Printf("joinRoom: %v", participantView)

		if participantView.SID == "" {
			return c.String(http.StatusBadRequest, "SID required")
		}

		var room Room
		db.Where("s_id=?", participantView.SID).First(&room)
		if room.SID != participantView.SID {
			return c.String(http.StatusNotFound, "")
		}

		var participants []Participant
		db.Where("s_id=?", room.SID).Find(&participants)

		var host Participant
		db.Where("uid=?", room.HostUID).First(&host)

		roomView := &RoomView{
			Title:    room.Title,
			Count:    len(participants),
			HostName: host.Name,
			E2EE:     room.E2EE,
		}

		return c.JSON(http.StatusOK, roomView)
	}
}

func callbackRoom(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
