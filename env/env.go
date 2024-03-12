package env

import (
	"fmt"
	"github.com/khivuksergey/portmonetka.authorization/errors"
	"os"
)

var (
	JwtSecret []byte
	JwtIssuer string

	PgUser     string
	PgPassword string
	PgDbName   string
	PgHost     string
	PgTimezone string
)

func LoadEnv() {
	var errMsg errors.ErrorMessage

	// JWT
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	if len(JwtSecret) == 0 {
		errMsg.Append(envErrorMsg("JWT_SECRET"))
	}
	JwtIssuer = os.Getenv("JWT_ISSUER")
	if JwtIssuer == "" {
		errMsg.Append(envErrorMsg("JWT_ISSUER"))
	}

	// Postgres
	PgUser = os.Getenv("POSTGRES_USER")
	if PgUser == "" {
		errMsg.Append(envErrorMsg("POSTGRES_USER"))
	}
	PgPassword = os.Getenv("POSTGRES_PASSWORD")
	if PgPassword == "" {
		errMsg.Append(envErrorMsg("POSTGRES_PASSWORD"))
	}
	PgDbName = os.Getenv("POSTGRES_DB_NAME")
	if PgDbName == "" {
		errMsg.Append(envErrorMsg("POSTGRES_DB_NAME"))
	}
	PgHost = os.Getenv("POSTGRES_HOST")
	if PgHost == "" {
		errMsg.Append(envErrorMsg("POSTGRES_HOST"))
	}
	PgTimezone = os.Getenv("POSTGRES_TIMEZONE")
	if PgTimezone == "" {
		errMsg.Append(envErrorMsg("POSTGRES_TIMEZONE"))
	}

	if err := errMsg.ToError(); err != nil {
		panic(err)
	}
}

func envErrorMsg(env string) string { return fmt.Sprintf("%s env variable missing", env) }
