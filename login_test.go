// login_service_test.go
package main

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		username string
		password string
		wantErr  bool
	}{
		{"admin", "password123", false},
		{"admin", "wrongpassword", true},
		{"wronguser", "password123", true},
		{"", "", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s/%s", tt.username, tt.password), func(t *testing.T) {
			err := Login(tt.username, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
