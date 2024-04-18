package handlers

import (
	"mini-project-3/entity"
	"mini-project-3/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BookingService struct {
	db *gorm.DB
}

func NewBookingService(db *gorm.DB) *BookingService {
	return &BookingService{db: db}
}

func (bs *BookingService) BookEquipment(c echo.Context) error {
	// Check if context is nil
	if c == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "context is nil")
	}

	// Validasi token JWT
	userID := helpers.GetUserId(c)

	// Check if userID is zero
	if userID == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "user ID not found in context")
	}

	// Bind request data to booking struct
	var booking entity.Booking
	if err := c.Bind(&booking); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	// Simpan pemesanan ke basis data
	booking.UserID = userID
	if err := bs.db.Create(&booking).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, booking)
}
