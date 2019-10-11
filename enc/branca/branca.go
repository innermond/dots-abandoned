package branca

import (
	"time"

	"github.com/hako/branca"
	"github.com/innermond/dots/enc"
)

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
func NewEncrypt(key string, ttl time.Duration) enc.Tokenizer {
	bk := branca.NewBranca(key)
	bk.SetTTL(uint32(ttl))
	return &Encrypt{b: bk}
}
