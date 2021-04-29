package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		ID:       1,
		Email:    "invalid_user@example.ru",
		Password: "secret",
	}
}
