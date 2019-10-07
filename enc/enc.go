package enc

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Tokenizer have capability to tokenize
type Tokenizer interface {
	Encode(string) (string, error)
	Decode(string) (string, error)
}

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

// RandKey return a uuid string
func RandKey() (string, error) {
	apikey, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", apikey), nil
}
