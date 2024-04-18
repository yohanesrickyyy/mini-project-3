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
		fmt.Println("Error creating activity log:", err)
		return err
	}
	fmt.Println("Activity log created successfully")
	return nil
}

func (us *UserService) TopUp(c echo.Context) error {
	// Retrieve user_id from the context and ensure it's an integer
	userID, ok := c.Get("user_id").(int)
	if !ok {
		fmt.Println("Invalid user ID format")
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid user ID format")
	}
	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := c.Bind(&payload); err != nil {
		fmt.Println("Error binding request payload:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	var user entity.User
	if err := us.db.First(&user, userID).Error; err != nil {
		fmt.Println("Error fetching user data:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch user data")
	}

	user.DepositAmount += payload.Amount

	if err := us.db.Save(&user).Error; err != nil {
		fmt.Println("Error saving user data:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to save user data")
	}

	desc := fmt.Sprintf("user with ID [%d] topped up deposit amount by %f", userID, payload.Amount)

	if err := us.CreateActivityLog(userID, desc); err != nil {
		fmt.Println("Error creating activity log:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create activity log")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "deposit topped up successfully",
		"user_id":      userID,
		"new_deposit":  user.DepositAmount,
		"topup_amount": payload.Amount,
	})
}
