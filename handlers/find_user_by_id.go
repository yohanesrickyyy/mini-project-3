package handlers

import (
	"mini-project-3/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (us *UserService) FindUserByID(c echo.Context) error {
	// Get user ID from the URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	// Find the user by ID in the database
	var user entity.User
	if err := us.db.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, user)
}
