package enc

import (
	"fmt"
	"sync"
	"time"

	"github.com/hako/branca"
	"github.com/innermond/dots/env"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Tokenizer have capability to tokenize
type Tokenizer interface {
	Encode(string) (string, error)
	Decode(string) (string, error)
}

// Encrypt adapt branca to Tokenizer interface
type Encrypt struct{ b *branca.Branca }

// Encode encode and encrypt a string
func (e *Encrypt) Encode(str string) (string, error) {
	return e.b.EncodeToString(str)
}

// Decode decode an encrypted string
func (e *Encrypt) Decode(str string) (string, error) {
	return e.b.DecodeToString(str)
}

// NewEncrypt adapts branca to Tokenizer interface
func NewEncrypt(key string, ttl time.Duration) Tokenizer {
	bk := branca.NewBranca(key)
	bk.SetTTL(uint32(ttl))
	return &Encrypt{b: bk}
}

var tok Tokenizer

func Tok() Tokenizer {
	return tok
}

func tokenizer() {
	tok = NewEncrypt(env.TokenKey(), time.Second*10)
}

func run() {
	tokenizer()
}

var once sync.Once

func Init() {
	once.Do(run)
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
func RandKey() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}
