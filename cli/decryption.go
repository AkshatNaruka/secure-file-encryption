package cli

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io/ioutil"
)

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
		return errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, plaintext, 0644)
}