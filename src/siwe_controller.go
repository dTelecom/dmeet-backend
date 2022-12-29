package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid/v4"
	"github.com/spruceid/siwe-go"
	"net/http"
)

// Nonce model
type Nonce struct {
	gorm.Model
	Nonce   string `gorm:"uniqueIndex"`
	Address string
}

// VerifyNonceRequest json
type VerifyNonceRequest struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
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

		var verifyNonceRequest VerifyNonceRequest
		err := c.Bind(&verifyNonceRequest)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		message, err := siwe.ParseMessage(verifyNonceRequest.Message)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var nonce Nonce
		db.Where("nonce=?", message.GetNonce()).First(&nonce)
		if nonce.Nonce != message.GetNonce() {
			return c.String(http.StatusNotFound, "")
		}

		if nonce.Address != "" {
			return c.String(http.StatusBadRequest, "used nonce")
		}

		_, err = message.VerifyEIP191(verifyNonceRequest.Signature)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		nonce.Address = message.GetAddress().String()
		db.Save(&nonce)

		return c.String(http.StatusOK, "")
	}
}
