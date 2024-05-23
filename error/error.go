package error

import (
	"errors"
	"fmt"
)

var (
	EmptyNamePassword            = errors.New("name and password cannot be empty")
	InvalidPassword              = errors.New("invalid password")
	TokenClaimsFail              = errors.New("failed to get token claims")
	TokenSignFail                = errors.New("failed to sign token")
	UserAlreadyExists            = errors.New("user with this name already exists")
	UserNotFound                 = errors.New("user was not found")
	GetTokenError                = func(err error) error { return fmt.Errorf("get token error: %w", err) }
	UpdateLastLoginTimeTimeError = func(err error) error { return fmt.Errorf("could not update last login time: %w", err) }
)

const (
	InvalidInputData     = "invalid input data"
	LoginFailed          = "login failed"
	CannotCreateUser     = "cannot create user"
	CannotUpdateUsername = "cannot update username"
	CannotUpdatePassword = "cannot update password"
	CannotDeleteUser     = "cannot delete user"
)

type ErrorMessage string

func (m *ErrorMessage) Append(errMessage string) {
	if *m != "" {
		*m += "; "
	}
	*m += ErrorMessage(errMessage)
}

func (m *ErrorMessage) ToError() error {
	if *m == "" {
		return nil
	}
	return fmt.Errorf(fmt.Sprint(*m))
}
