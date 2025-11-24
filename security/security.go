package security

import "golang.org/x/crypto/bcrypt"

// HashPassword converts a plain text password into a secure hash.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword compares a hashed password with the login password.
func CheckPassword(hashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
