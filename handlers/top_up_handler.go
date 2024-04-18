package handlers

import (
	"fmt"
	"mini-project-3/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (us *UserService) CreateActivityLog(userID int, description string) error {
	activityLog := entity.UserActivity{
		UserId:      userID,
		Description: description,
	}
	if err := us.db.Create(&activityLog).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) TopUp(c echo.Context) error {
	userID := int(c.Get("user_id").(float64))

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	var user entity.User
	if err := us.db.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user.DepositAmount += payload.Amount

	if err := us.db.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	desc := fmt.Sprintf("user with ID [%d] topped up deposit amount by %f", userID, payload.Amount)

	if err := us.CreateActivityLog(userID, desc); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "deposit topped up successfully",
		"user_id":      userID,
		"new_deposit":  user.DepositAmount,
		"topup_amount": payload.Amount,
	})
}
