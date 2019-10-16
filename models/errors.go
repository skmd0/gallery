package models

import "strings"

const (
	// ErrNotFound is returned when a resource cannot be found in the database
	ErrNotFound modelError = "models: resource not found"

	// ErrPasswordIncorrect is returned when an invalid password
	// is used when attenpting to authenticate a user.
	ErrPasswordIncorrect modelError = "models: incorrect password provided"

	// ErrEmailRequired is returned when en email address is not provided when creating a user
	ErrEmailRequired modelError = "models: email address is required"

	// ErrEmailInvalid is returned when an email address is not provided
	ErrEmailInvalid modelError = "models: email address is invalid"

	// ErrEmailAlreadyTaken is returned when email is already registered in DB
	ErrEmailAlreadyTaken modelError = "models: email address is already taken"

	// ErrPasswordTooShort is returned when password is too short
	ErrPasswordTooShort modelError = "models: password must be at least 8 characters long"

	// ErrPasswordRequired is returned when password is not set
	ErrPasswordRequired modelError = "models: password is required"

	// ErrTitleRequired when gallery title is not provided
	ErrTitleRequired modelError = "models: title is required"

	// ErrTokenInvalid when password reset token is not valid
	ErrTokenInvalid modelError = "models: token provided is not valid"

	// ErrIDInvalid is returned when an invalid ID is provided
	ErrIDInvalid privateError = "models: ID provided was invalid"

	// ErrRememberTooShort is returned when the remember token is less than 32 bytes
	ErrRememberTooShort privateError = "models: remeber token must be at least 32 bytes"

	// ErrRememberRequired is returned when create or update is attempted
	// without a user remeber token hash
	ErrRememberRequired privateError = "models: remember token hash is required"

	// ErrUserIDRequired requires id
	ErrUserIDRequired privateError = "models: user ID is required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	return strings.Title(s)
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}
