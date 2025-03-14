package sdk

import "github.com/google/uuid"

// GenerateUniqueID generates a unique identifier string of 8 characters in length.
// It uses the UUID package to create a new UUID and returns the first 8 characters of the UUID string.
func GenerateUniqueID() string {
	return uuid.New().String()[:8]
}
