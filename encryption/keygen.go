package cli

import (
	
	"crypto/rand"
	
)

// GenerateAESKey generates a secure AES encryption key.
func GenerateAESKey() ([]byte, error) {
	key := make([]byte, 32) // 32 bytes for AES-256
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
