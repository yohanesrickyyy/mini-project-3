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

// BookEquipment godoc
// @Summary Book equipment
// @Description Book equipment by providing booking details
// @Tags bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "JWT token"
// @Param booking body entity.Booking{} true "Booking details"
// @Success 201 {object} Booking "Successfully booked equipment"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 500 {object} HTTPError "Internal server error"
// @Router /bookings [post]
func (bs *BookingService) BookEquipment(c echo.Context) error {
	if c == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "context is nil")
	}

	userID := helpers.GetUserId(c)

	if userID == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "user ID not found in context")
	}

	var booking entity.Booking
	if err := c.Bind(&booking); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	booking.UserID = userID
	if err := bs.db.Create(&booking).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, booking)
}
