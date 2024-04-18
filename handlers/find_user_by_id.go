package handlers

import (
	"mini-project-3/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// FindUserByID godoc
// @Summary Find a user by ID
// @Description Get a user's details by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} entity.User
// @Router /users/{id} [get]
func (us *UserService) FindUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	var user entity.User
	if err := us.db.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, user)
}
