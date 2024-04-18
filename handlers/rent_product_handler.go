package handlers

import (
	"mini-project-3/entity"
	"mini-project-3/helpers"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RentProduct struct {
	db *gorm.DB
}

func RentProductService(db *gorm.DB) *RentProduct {
	return &RentProduct{db: db}
}

func (as *RentProduct) RentEquipmentHandler(c echo.Context) error {
	userID := helpers.GetUserId(c)

	var request struct {
		EquipmentID int       `json:"equipment_id"`
		RentalDate  time.Time `json:"rental_date"`
		ReturnDate  time.Time `json:"return_date"`
	}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	// Periksa apakah peralatan tersedia
	var equipment entity.EquipmentType
	if err := as.db.First(&equipment, request.EquipmentID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "equipment not found")
	}
	if !equipment.Availability {
		return echo.NewHTTPError(http.StatusBadRequest, "equipment is not available")
	}

	// Buat entri transaksi baru
	transaction := entity.Transaction{
		RentalID:        0,
		TransactionDate: time.Now(),
		Amount:          0,
		PaymentMethod:   "Cash",
	}
	if err := as.db.Create(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Buat entri rental baru
	rental := entity.EquipmentRentalHistory{
		UserID:      userID,
		EquipmentID: uint(request.EquipmentID),
		RentalDate:  request.RentalDate,
		ReturnDate:  request.ReturnDate,
	}
	if err := as.db.Create(&rental).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	equipment.Availability = false
	if err := as.db.Save(&equipment).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":        "equipment rented successfully",
		"rental_id":      rental.RentalID,
		"transaction_id": transaction.TransactionID,
	})
}
