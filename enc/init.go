package enc

import (
	"sync"
	"time"

	"github.com/innermond/dots/env"
)

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
