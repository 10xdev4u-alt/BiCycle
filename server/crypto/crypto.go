package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/config"
)

// Decrypt AES-256-CBC encrypted text
func Decrypt(encryptedText string, cfg *config.Config) (string, error) {
	key := []byte(cfg.JWTSecret) // Using JWTSecret for now, should be a dedicated key
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad
	padding := int(ciphertext[len(ciphertext)-1])
	if padding > len(ciphertext) {
		return "", errors.New("invalid padding")
	}
	plaintext := ciphertext[:len(ciphertext)-padding]

	return string(plaintext), nil
}

// Encrypt encrypts text using AES-256-CBC
func Encrypt(plaintext string, cfg *config.Config) (string, error) {
    key := []byte(cfg.JWTSecret)
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    // Pad plaintext to be a multiple of the block size
    padding := aes.BlockSize - len(plaintext)%aes.BlockSize
    padtext := append([]byte(plaintext), bytes.Repeat([]byte{byte(padding)}, padding)...)

    ciphertext := make([]byte, aes.BlockSize+len(padtext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], padtext)

    return base64.StdEncoding.EncodeToString(ciphertext), nil
}
