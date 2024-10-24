package services

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "strongpassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Errorf("expected hashed password, got empty string")
	}

	// Test with weak password (less than 8 characters)
	_, err = HashPassword("short")
	if err == nil {
		t.Errorf("expected an error for short password, got nil")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "strongpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Correct password
	err := checkPassword(string(hashedPassword), password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Incorrect password
	err = checkPassword(string(hashedPassword), "wrongpassword")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}
