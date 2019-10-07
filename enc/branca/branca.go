package branca

import (
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
func NewEncrypt(bkey []byte, ttl uint32) enc.Tokenizer {
	bk := branca.NewBranca(string(bkey))
	bk.SetTTL(ttl)
	return &Encrypt{b: bk}
}
