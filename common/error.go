package common

import (
	"errors"
	"fmt"
)

var (
	EmptyNamePassword = errors.New("name and password cannot be empty")
	InvalidPassword   = errors.New("invalid password")
	InvalidUserData   = errors.New("invalid user data")
	NilUserToken      = errors.New("cannot get token for nil user")
	TokenClaimsFail   = errors.New("failed to get token claims")
	UserAlreadyExists = errors.New("user with this name already exists")
	UserNotFound      = errors.New("user was not found")
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
