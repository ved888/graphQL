package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

// HashString generates SHA256 for a given string
//func HashString(toHash string) string {
//	sha := sha512.New()
//	sha.Write([]byte(toHash))
//	return hex.EncodeToString(sha.Sum(nil))
//}

// RequestErr models contains the body having details related with some kind of error
// which happened during processing of a request
type RequestErr struct {
	// ID for the request
	// Example: 8YeCqPXmM
	ID string `json:"id"`

	// MessageToUser will contain error message
	// Example: Invalid Email
	MessageToUser string `json:"messageToUser"`

	// DeveloperInfo will contain additional developer info related with error
	// Example: Invalid email format
	DeveloperInfo string `json:"developerInfo"`

	// Err contains the error or exception message
	// Example: validation on email failed with error invalid email format
	Err string `json:"error"`

	// StatusCode will contain the status code for the error
	// Example: 500
	StatusCode int `json:"statusCode"`

	// IsClientError will be false if some internal server error occurred
	IsClientError bool `json:"isClientError"`
} // @name RequestErr

type Branch string

const (
	Production  Branch = "main"
	Staging     Branch = "stage"
	Development Branch = "dev"
)

// HashString generates SHA256 for a given string
func HashString(toHash string) string {
	sha := sha512.New()
	sha.Write([]byte(toHash))
	return hex.EncodeToString(sha.Sum(nil))
}

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// IsProd returns true if running on prod
func IsProd() bool {
	return GetBranch() == Production
}

// GetBranch returns current branch name, defaults to development if no branch specified
func GetBranch() Branch {
	b := os.Getenv("BRANCH")
	if b == "" {
		return Development
	}
	return Branch(b)
}

// IsBranchEnvSet checks if the branch environment is set
func IsBranchEnvSet() bool {
	b := os.Getenv("BRANCH")
	return b != ""
}
