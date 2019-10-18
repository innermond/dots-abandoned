package enc

import (
	"golang.org/x/crypto/bcrypt"
)

// Password generate an encrypted password
func Password(pwd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// HashIsPassword can a hashed string be the plain one if the later would be hashed
func HashIsPassword(hashed string, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))
}
