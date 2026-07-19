package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"runtime"
)

var (
	ErrEncrypt       = errors.New("encryption failed")
	ErrDecrypt       = errors.New("decryption failed")
	encryptionKeyDir = ""
)

func SetEncryptionKeyDir(dir string) {
	encryptionKeyDir = dir
}

func deriveKey() []byte {
	keyFile := ""
	if encryptionKeyDir != "" {
		keyFile = encryptionKeyDir + "/termcode.key"
	}

	if keyFile != "" {
		if data, err := os.ReadFile(keyFile); err == nil && len(data) == 32 {
			return data
		}
	}

	hostname, _ := os.Hostname()
	seed := hostname + "-" + runtime.GOARCH + "-" + runtime.GOOS
	hash := sha256.Sum256([]byte(seed))
	key := hash[:]

	if keyFile != "" {
		if err := os.MkdirAll(encryptionKeyDir, 0o700); err == nil {
			os.WriteFile(keyFile, key, 0o600)
		}
	}

	return key
}

func EncryptAPIKey(plaintext string) (string, error) {
	key := deriveKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", ErrEncrypt
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", ErrEncrypt
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", ErrEncrypt
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nil
}

func DecryptAPIKey(encoded string) (string, error) {
	key := deriveKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", ErrDecrypt
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", ErrDecrypt
	}

	ciphertext, err := hex.DecodeString(encoded)
	if err != nil {
		return "", ErrDecrypt
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", ErrDecrypt
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", ErrDecrypt
	}

	return string(plaintext), nil
}

func MaskAPIKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "****" + key[len(key)-4:]
}
