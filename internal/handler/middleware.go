package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/khivuksergey/webserver/logger"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

var (
	InvalidToken       = echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	UnauthorizedAccess = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized access")
	InvalidPathParam   = echo.NewHTTPError(http.StatusUnauthorized, "invalid path param userId")
)

type Middleware struct {
	logger logger.Logger
	JWT    echo.MiddlewareFunc
}

func NewMiddleware(logger logger.Logger) *Middleware {
	return &Middleware{
		logger: logger,
		JWT:    echojwt.JWT([]byte(viper.GetString("JWT_SECRET"))),
	}
}

// Authentication middleware checks if path param "userId" is the same as the subject in JWT from the Context
func (m *Middleware) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
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
