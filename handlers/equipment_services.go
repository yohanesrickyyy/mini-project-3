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

func (es *EquipmentService) PostEquipmentHandler(c echo.Context) error {
	// Bind request payload to Equipment struct
	var equipment entity.EquipmentType
	if err := c.Bind(&equipment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	// Set availability to true for new equipment
	equipment.Availability = true

	// Set default rental costs and category if not provided
	if equipment.RentalCosts == 0 {
		equipment.RentalCosts = 0.0
	}
	if equipment.Category == "" {
		equipment.Category = "General"
	}

	// Create equipment in the database
	if err := es.db.Create(&equipment).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, equipment)
}

func (es *EquipmentService) GetAllEquipmentHandler(c echo.Context) error {
	// Fetch all equipment from the database
	var equipment []entity.EquipmentType
	if err := es.db.Find(&equipment).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, equipment)
}
