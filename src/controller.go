package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Participant model
type Participant struct {
	gorm.Model
	Name      string
	UID       string
	SID       string
	IsHost    bool
	AddedAt   time.Time
	RemovedAt time.Time
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
	Count            int64  `json:"count"`
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
	CallID   string `json:"callID"`
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
	NodePK string
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

		url, nodeID, nodePK, err := getNodeURL()
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

		tokenString, signature, err := GetTokenSignature(token)
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
			NodePK: nodePK,
		}
		db.Create(&call)

		participant := &Participant{
			Name:      roomRequest.Name,
			SID:       SID,
			UID:       UID,
			IsHost:    true,
			AddedAt:   time.Time{},
			RemovedAt: time.Time{},
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

		url, nodeID, nodePK, err := getNodeURL()
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
				NodePK: nodePK,
			}
			db.Create(&call)
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

		tokenString, signature, err := GetTokenSignature(token)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		participant := &Participant{
			Name:      roomRequest.Name,
			SID:       roomRequest.SID,
			UID:       UID,
			IsHost:    false,
			AddedAt:   time.Time{},
			RemovedAt: time.Time{},
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

		var count int64
		db.Model(&Participant{}).Where("s_id=? AND added_at!=? AND removed_at=?", room.SID, time.Time{}, time.Time{}).Count(&count)

		var host Participant
		db.Where("uid=?", room.HostUID).First(&host)

		roomView := &RoomView{
			Title:            room.Title,
			Count:            count,
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

		var notifyData NotifyData
		err = json.Unmarshal(notifyRequest.Data, &notifyData)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		log.Printf("got data: %v", notifyData)

		var call Call
		db.Where("call_id=?", notifyData.CallID).First(&call)
		if call.CallID != notifyData.CallID {
			return c.String(http.StatusNotFound, "")
		}

		verified, err := VerifySignature(notifyRequest.Data, call.NodePK, notifyRequest.Signature)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if verified != true {
			return c.String(http.StatusBadRequest, "not verified signature")
		}

		var participant Participant
		db.Where("uid=?", notifyData.UID).First(&participant)
		if participant.UID != notifyData.UID {
			return c.String(http.StatusNotFound, "")
		}

		if notifyData.Type == "join" {
			participant.AddedAt = time.Now()
		}
		if notifyData.Type == "leave" {
			participant.RemovedAt = time.Now()
		}

		db.Save(&participant)

		epoch, err := getEpochHeight()
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		message := notifyData.CallID + ":" + strconv.Itoa(notifyData.Duration) + ":" + strconv.Itoa(int(epoch))
		sig, err := GetSignature([]byte(message))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		notifyResponse := NotifyResponse{
			Signature: sig,
			Epoch:     epoch,
		}

		return c.JSON(http.StatusOK, notifyResponse)
	}
}
