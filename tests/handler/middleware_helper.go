package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/khivuksergey/portmonetka.authorization/internal/handler"
	"net/http"
	"strconv"
)

const userId = uint64(111)

var (
	userIdStr    = strconv.FormatUint(userId, 10)
	invalidIdStr = userIdStr + "1"
)

type testCase struct {
	token              *jwt.Token
	pathParamValue     string
	expectedError      error
	expectedStatusCode int
}

var testCases = []testCase{
	{
		token:              nil,
		pathParamValue:     "",
		expectedError:      handler.InvalidToken,
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		token:              jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}),
		pathParamValue:     "",
		expectedError:      handler.InvalidToken,
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		token:              jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": userId}),
		pathParamValue:     "",
		expectedError:      handler.InvalidPathParam,
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		token:              jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": userId}),
		pathParamValue:     invalidIdStr,
		expectedError:      handler.UnauthorizedAccess,
		expectedStatusCode: http.StatusUnauthorized,
	},
}
