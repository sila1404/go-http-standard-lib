package auth

import (
	"testing"
)

func TestHashedPassword(t *testing.T) {
	hash, err := HashedPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to be not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashedPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, []byte("password")) {
		t.Errorf("expected password to match hash")
	}
	if ComparePassword(hash, []byte("notpassword")) {
		t.Errorf("expected password to not match hash")
	}
}
