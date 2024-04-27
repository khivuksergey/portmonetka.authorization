package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/webserver/logger"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"strconv"
)

const (
	invalidToken       = "invalid token"
	invalidPathParam   = "invalid path param userId"
	unauthorizedAccess = "unauthorized access"
)

type AuthenticationMiddleware struct {
	logger logger.Logger
	jwt    echo.MiddlewareFunc
}

func NewAuthenticationMiddleware(logger logger.Logger) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		logger: logger,
		jwt:    echojwt.JWT([]byte(viper.GetString("JWT_SECRET"))),
	}
}

func (a *AuthenticationMiddleware) JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return a.jwt(next)
}

// Authenticate checks if path param "userId" is the same as the subject in JWT from the Context
func (a *AuthenticationMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// take user token from jwt
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return common.NewAuthorizationError(invalidToken, common.NoUserToken)
		}

		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok {
			return common.NewAuthorizationError(invalidToken, common.InvalidTokenClaims)
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			return common.NewAuthorizationError(invalidToken, common.InvalidSubjectClaim)
		}
		subject := uint64(sub)

		// take userId path param
		userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
		if err != nil {
			return common.NewAuthorizationError(invalidPathParam, nil)
		}

		if subject != userId {
			return common.NewAuthorizationError(unauthorizedAccess, nil)
		}

		c.Set("userId", userId)

		return next(c)
	}
}

func (a *AuthenticationMiddleware) AuthenticateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return a.JWT(a.Authenticate(next))
}
