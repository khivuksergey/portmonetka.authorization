package common

import (
	"errors"
	"fmt"
)

var (
	InvalidUserData   = errors.New("invalid user data")
	EmptyName         = errors.New("name cannot be empty")
	EmptyPassword     = errors.New("password cannot be empty")
	EmptyNamePassword = errors.New("name and password cannot be empty")
	UserAlreadyExists = errors.New("user with this name already exists")
	NilUserToken      = errors.New("cannot get token for nil user")
	TokenClaimsFail   = errors.New("failed to get token claims")
	UserNotFound      = errors.New("user was not found")
	InvalidPassword   = errors.New("invalid password")
	GetTokenError     = func(err error) error { return fmt.Errorf("get token error: %v", err) }
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
