package handlers

import (
	"errors"
	"mini-project-3/dto"
	"mini-project-3/entity"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Login struct {
	db *gorm.DB
}

func NewLogin(db *gorm.DB) *Login {
	return &Login{db: db}
}

// LoginHandler godoc
// @Summary Login user
// @Description Log in with email and password to get authentication token
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} HTTPError
// @Router /login [post]
func (as *Login) LoginHandler(c echo.Context) error {
	login := new(dto.Login)

	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := new(entity.User)
	res := as.db.Where("email = ?", login.Email).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "email not found")
	}
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "incorrect password")
	}

	tokenString, err := CreateJWT(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   tokenString,
	})
}

func CreateJWT(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "failed to create token: ", err
	}

	return tokenString, nil
}
