package repository

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

type (
	accessKeyAesSha256Cbc struct {
		appKey string
	}
)

func NewAccessKeyAesSha256Cbc(appKey string) AccessKey {
	return &accessKeyAesSha256Cbc{appKey: appKey}
}

func (a accessKeyAesSha256Cbc) Encrypt(currTime time.Time) string {
	timestamp := currTime.Unix()
	plain := fmt.Sprintf("%s@%d", a.appKey, timestamp)

	// PHP serialize
	serialized := fmt.Sprintf("s:%d:\"%s\";", len(plain), plain)

	// Key and IV: use hex string of SHA-256 hash
	keyBytes := sha256.Sum256([]byte(a.appKey))
	keyHex := hex.EncodeToString(keyBytes[:]) // 64 chars

	key := []byte(keyHex[:32]) // first 32 bytes of hex string
	iv := []byte(keyHex[:16])  // first 16 bytes of hex string

	// PKCS7 padding
	padded := pkcs7Pad([]byte(serialized), aes.BlockSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, padded)

	encodeToString := base64.StdEncoding.EncodeToString(ciphertext)
	return base64.StdEncoding.EncodeToString([]byte(encodeToString))
}

func (a accessKeyAesSha256Cbc) Decrypt(encrypted string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(string(decodeString))
	if err != nil {
		return "", err
	}

	// Reconstruct key and IV from appKey
	keyBytes := sha256.Sum256([]byte(a.appKey))
	keyHex := hex.EncodeToString(keyBytes[:])

	key := []byte(keyHex[:32]) // first 32 chars of hex string
	iv := []byte(keyHex[:16])  // first 16 chars of hex string

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("cipher creation failed: %w", err)
	}

	// Decrypt
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	// Remove PKCS7 padding
	unpadded, err := pkcs7Unpad(decrypted, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf("unpad failed: %w", err)
	}

	return string(unpadded), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, fmt.Errorf("invalid padded data")
	}
	padLen := int(data[len(data)-1])
	if padLen == 0 || padLen > blockSize {
		return nil, fmt.Errorf("invalid padding length")
	}
	for _, b := range data[len(data)-padLen:] {
		if int(b) != padLen {
			return nil, fmt.Errorf("invalid padding byte")
		}
	}
	return data[:len(data)-padLen], nil
}
