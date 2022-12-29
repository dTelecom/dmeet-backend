package main

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/zerolog/log"
	"github.com/spruceid/siwe-go"
	"net/http"
)

// Nonce model
type Nonce struct {
	gorm.Model
	Nonce   string `gorm:"uniqueIndex"`
	Address string
}

func getNonce(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

		nonce := &Nonce{
			Nonce: shortuuid.New(),
		}
		db.Create(&nonce)

		return c.String(http.StatusOK, nonce.Nonce)
	}
}

func verifyNonce(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request().Body)
		raw := buf.String()
		log.Printf("got: %v", raw)

		message, err := siwe.ParseMessage(raw)
		log.Printf("parsed: %v %v", message, err)
		return c.String(http.StatusBadRequest, "")
	}
}
