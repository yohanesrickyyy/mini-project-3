package helpers

import (
	"mini-project-3/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetToken(tokenString string) (*jwt.Token, error) {
	claims := &config.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTKey(), nil
	})
	return token, err
}

func GetUserId(c echo.Context) int {
	tokenCookie, _ := c.Cookie("token")
	tokenString := tokenCookie.Value
	token, _ := GetToken(tokenString)
	claims := token.Claims.(*config.JWTClaim)
	return claims.UserID
}
