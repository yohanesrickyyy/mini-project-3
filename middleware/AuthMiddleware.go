package middlewares

import (
	"mini-project-3/helpers"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenCookie, err := c.Cookie("token")
			if err != nil {
				errResponse := helpers.ErrUnauthorized(err.Error())
				return c.JSON(errResponse.Code, errResponse)
			}
			tokenString := tokenCookie.Value
			token, err := helpers.GetToken(tokenString)

			if err != nil || !token.Valid {
				errResponse := helpers.ErrUnauthorized(err.Error())
				return c.JSON(errResponse.Code, errResponse)
			}
			return next(c)
		}
	}
}
