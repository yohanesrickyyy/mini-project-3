package handlers

import (
	"mini-project-3/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Registration struct {
	db *gorm.DB
}

func NewRegistration(db *gorm.DB) *Registration {
	return &Registration{db: db}
}

func (as *Registration) RegisterHandler(c echo.Context) error {
	user := new(entity.User)

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(hashedPassword)

	if err := as.db.Create(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}
