package middleware

import (
	"github.com/labstack/echo-jwt/v4"
	"os"
)

var JWTAuthorization = echojwt.JWT([]byte(os.Getenv("JWT_SECRET")))
