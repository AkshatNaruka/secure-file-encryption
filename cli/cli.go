package cli

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
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

// EncryptFile encrypts a file using AES-GCM.
func EncryptFile(inputPath, outputPath string, key []byte) error {
	plaintext, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ioutil.WriteFile(outputPath, ciphertext, 0644)
}

// DecryptFile decrypts a file using AES-GCM.
func DecryptFile(inputPath, outputPath string, key []byte) error {
	ciphertext, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil // Handle this case as needed
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, plaintext, 0644)
}
