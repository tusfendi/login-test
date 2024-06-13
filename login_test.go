// login_service_test.go
package main

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		wantErr  bool
		err      error
	}{
		{
			name:     "Error username Required",
			username: "",
			password: "password123",
			wantErr:  true,
			err:      errors.New("username is required"),
		},
		{
			name:     "Error password Required",
			username: "admin",
			password: "",
			wantErr:  true,
			err:      errors.New("password is required"),
		},
		{
			name:     "Success Login",
			username: "admin",
			password: "password123",
			wantErr:  false,
		},
		{
			name:     "Error invalid creds",
			username: "admin",
			password: "1289u",
			wantErr:  true,
			err:      errors.New("invalid credentials"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Login(tt.username, tt.password)
			if (err != nil) != tt.wantErr || !assert.IsEqual(tt.err, err) {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
