package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// User represents a user in the system
type User struct {
	Username     string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Salt         string `json:"salt"`
}

// GenerateSalt generates a random salt
func GenerateSalt() (string, error) {
	salt := make([]byte, 16) // 16 bytes of salt
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes the password with the given salt
func HashPassword(password, salt string) string {
	saltedPassword := password + salt
	hash := sha256.Sum256([]byte(saltedPassword))
	return base64.StdEncoding.EncodeToString(hash[:])
}

// CreateUser creates a new user with a hashed password and salt
func CreateUser(username, password string) (*User, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword := HashPassword(password, salt)

	user := &User{
		Username:     username,
		HashedPassword: hashedPassword,
		Salt:         salt,
	}

	return user, nil
}

// VerifyPassword checks if the provided password matches the stored hashed password
func (u *User) VerifyPassword(password string) bool {
	hashedInputPassword := HashPassword(password, u.Salt)
	return hashedInputPassword == u.HashedPassword
}

func main() {
	// Example usage
	username := "exampleUser"
	password := "securePassword123"

	user, err := CreateUser(username, password)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	fmt.Printf("User created: %+v\n", user)

	// Verify password
	isValid := user.VerifyPassword("securePassword123")
	fmt.Println("Password valid:", isValid)

	isValid = user.VerifyPassword("wrongPassword")
	fmt.Println("Password valid:", isValid)
}

