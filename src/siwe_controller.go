package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid/v4"
	"net/http"
)

// Nonce model
type Nonce struct {
	gorm.Model
	Nonce string `gorm:"uniqueIndex"`
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
