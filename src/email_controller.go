package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Email model
type Email struct {
	gorm.Model
	Email string `gorm:"uniqueIndex"`
}

// EmailRequest model
type EmailRequest struct {
	Email string `json:"email"`
}

func saveEmail(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var emailRequest EmailRequest
		err := c.Bind(&emailRequest)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		log.Printf("saveEmail: %v", emailRequest)

		if emailRequest.Email == "" {
			return c.String(http.StatusBadRequest, "Email required")
		}

		email := &Email{
			Email: emailRequest.Email,
		}
		db.Create(&email)

		return c.JSON(http.StatusOK, "")
	}
}
