package handlers

import (
	"mini-project-3/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Registration struct {
	db *gorm.DB
}

func NewRegistration(db *gorm.DB) *Registration {
	return &Registration{db: db}
}

// RegisterHandler handles user registration.
// @Summary Register user
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body User true "User object"
// @Success 201 {object} User "Successfully registered user"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 500 {object} HTTPError "Internal server error"
// @Router /register [post]
func (as *Registration) RegisterHandler(c echo.Context) error {
	getInputUser := new(entity.User)

	// Bind data dari body permintaan ke objek userData
	if err := c.Bind(&getInputUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Buat objek user berdasarkan data yang diterima
	user := &entity.User{
		Email:         getInputUser.Email,
		Password:      getInputUser.Password,
		DepositAmount: getInputUser.DepositAmount,
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(hashedPassword)

	// Simpan user ke basis data
	if err := as.db.Create(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Hapus password dari respons JSON yang dikirimkan
	user.Password = ""

	// Kirim respons sukses dengan data pengguna yang didaftarkan
	return c.JSON(http.StatusCreated, user)
}
