package pwd

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const (
	passwordLength = 8
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
)

func GenerateRandomPassword() (string, error) {
	password := make([]byte, passwordLength)

	_, err := rand.Read(password)
	if err != nil {
		return "", err
	}

	for i := 0; i < passwordLength; i++ {
		password[i] = charset[int(password[i])%len(charset)]
	}

	return string(password), nil
}

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func Hash(password, salt string) string {
	saltedPassword := password + salt
	hash := sha256.Sum256([]byte(saltedPassword))
	return base64.StdEncoding.EncodeToString(hash[:])
}
