package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	InvalidToken       = echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	UnauthorizedAccess = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized access")
	InvalidPathParam   = echo.NewHTTPError(http.StatusUnauthorized, "invalid path param userId")
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// take user token from jwt
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return InvalidToken
		}

		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok {
			return InvalidToken
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			return InvalidToken
		}
		subject := uint64(sub)

		// take userId path param
		userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
		if err != nil {
			return InvalidPathParam
		}

		if subject != userId {
			return UnauthorizedAccess
		}

		c.Set("userId", userId)

		return next(c)
	}
}
