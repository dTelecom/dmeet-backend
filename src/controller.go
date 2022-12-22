package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Participant model
type Participant struct {
	gorm.Model
	Name string
	UID  string
	SID  string
}

// ParticipantView model
type ParticipantView struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
	SID  string `json:"sid"`
}

// Room model
type Room struct {
	gorm.Model
	SID      string
	CallID   string
	Password string
}

// RoomView model
type RoomView struct {
	SID      string `json:"sid"`
	UID      string `json:"uid"`
	URL      string `json:"url"`
	Password string `json:"password"`
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
		callID := shortuuid.New()
		password := generatePassword()

		url, err := getNodeURL()
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		room := &Room{
			SID:      SID,
			CallID:   callID,
			Password: password,
		}
		db.Create(&room)

		participant := &Participant{
			Name: participantView.Name,
			SID:  SID,
			UID:  UID,
		}
		db.Create(&participant)

		roomView := &RoomView{
			SID:      SID,
			UID:      UID,
			URL:      url,
			Password: password,
		}

		return c.JSON(http.StatusOK, roomView)
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

		participant := &Participant{
			Name: participantView.Name,
			SID:  participantView.SID,
			UID:  UID,
		}
		db.Create(&participant)

		roomView := &RoomView{
			SID:      participantView.SID,
			UID:      UID,
			URL:      url,
			Password: room.Password,
		}

		return c.JSON(http.StatusOK, roomView)
	}
}
