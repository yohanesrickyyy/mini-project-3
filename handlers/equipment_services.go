package handlers

import (
	"mini-project-3/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EquipmentService struct {
	db *gorm.DB
}

func NewEquipmentService(db *gorm.DB) *EquipmentService {
	return &EquipmentService{db: db}
}

// BookEquipment godoc
// @Summary Book equipment
// @Description Book equipment by providing booking details
// @Tags bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "JWT token"
// @Param booking body entity.Booking true "Booking details"
// @Success 201 {object} entity.Booking
// @Router /bookings [post]
func (es *EquipmentService) PostEquipmentHandler(c echo.Context) error {
	var equipment entity.EquipmentType
	if err := c.Bind(&equipment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	equipment.Availability = true

	if equipment.RentalCosts == 0 {
		equipment.RentalCosts = 0.0
	}
	if equipment.Category == "" {
		equipment.Category = "General"
	}

	if err := es.db.Create(&equipment).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, equipment)
}

func (es *EquipmentService) GetAllEquipmentHandler(c echo.Context) error {
	var equipment []entity.EquipmentType
	if err := es.db.Find(&equipment).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, equipment)
}
