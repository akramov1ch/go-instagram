package errors

import (
	"errors"
)

func ErrorEmail() error {
	return errors.New("Invalid email")
}

func HaveEmail() error {
	return errors.New("Email already exists")
}

func ErrorPassword() error {
	return errors.New("Password must be at least 8 characters long")
}

func HaveUsername() error {
	return errors.New("Username already exists")
}

func ErrorUsername() error {
	return errors.New("Invalid username")
}

func ErrorUserNotFound() error {
	return errors.New("User not found")
}