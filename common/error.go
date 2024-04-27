package common

import (
	"errors"
	"fmt"
)

var (
	NoUserToken                  = errors.New("failed to get user token")
	InvalidTokenClaims           = errors.New("invalid token claims")
	InvalidSubjectClaim          = errors.New("invalid subject claim")
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

type BaseError struct {
	Message string
	Cause   error
}

func (e BaseError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e BaseError) Unwrap() error {
	return e.Cause
}

type AuthorizationError struct{ BaseError }

func NewAuthorizationError(message string, cause error) AuthorizationError {
	return AuthorizationError{BaseError{Message: message, Cause: cause}}
}

type ValidationError struct{ BaseError }

func NewValidationError(message string, cause error) ValidationError {
	return ValidationError{BaseError{Message: message, Cause: cause}}
}

type UnprocessableEntityError struct{ BaseError }

func NewUnprocessableEntityError(message string, cause error) UnprocessableEntityError {
	return UnprocessableEntityError{BaseError{Message: message, Cause: cause}}
}

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
