package errors

import (
	"errors"
	"fmt"
)

var (
	UserNotFound            = errors.New("user was not found")
	InvalidPassword         = errors.New("invalid password")
	UserDataValidationError = errors.New("user data validation error")
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
