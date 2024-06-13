package main

import (
	"errors"
	"fmt"
)

// Hard-coded credentials
var (
	validUsername = "admin"
	validPassword = "password123"
)

func main() {
	// Example usage
	err := Login("admin", "password123")
	if err != nil {
		fmt.Println("Login failed:", err)
	} else {
		fmt.Println("Login successful")
	}
}

// Login function checks if the provided username and password are correct
func Login(username, password string) error {
	if username == "" {
		return errors.New("username is required")
	}

	if password == "" {
		return errors.New("password is required")
	}

	if username == validUsername && password == validPassword {
		return nil
	}

	return errors.New("invalid credentials")
}
