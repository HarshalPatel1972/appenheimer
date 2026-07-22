package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	// Stub for Argon2id hashing
	return "hashed_" + password, nil
}

func VerifyPassword(password, hash string) (bool, error) {
	// Stub for Argon2id verification
	return "hashed_"+password == hash, nil
}
