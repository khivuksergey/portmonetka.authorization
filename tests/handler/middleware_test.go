package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/khivuksergey/portmonetka.authorization/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthentication_Success(t *testing.T) {
	claims := jwt.MapClaims{
		"sub": userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token.Raw)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.Set("user", token)
	c.SetParamNames("userId")
	c.SetParamValues(userIdStr)

	middleware := handler.NewMiddleware(nil)
	authMiddleware := middleware.Authentication(func(c echo.Context) error {
		id, ok := c.Get("userId").(uint64)
		if !ok || id != userId {
			return c.String(http.StatusUnauthorized, "Invalid userId in context in next handler")
		}
		return c.String(http.StatusOK, "Success")
	})

	assert.NoError(t, authMiddleware(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Success", rec.Body.String())
}

func TestAuthentication_Errors(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	middleware := handler.NewMiddleware(nil)
	authMiddleware := middleware.Authentication(func(c echo.Context) error {
		id, ok := c.Get("userId").(uint64)
		if !ok || id != userId {
			return c.String(http.StatusUnauthorized, "Invalid userId in context in next handler")
		}
		return c.String(http.StatusOK, "Success")
	})

	for _, test := range testCases {
		if test.token != nil {
			req.Header.Set("Authorization", "Bearer "+test.token.Raw)
			c.Set("user", test.token)
		}
		c.SetParamNames("userId")
		c.SetParamValues(test.pathParamValue)

		assert.ErrorIs(t, authMiddleware(c), test.expectedError)
		assert.Equal(t, test.expectedStatusCode, rec.Code)
	}
}
