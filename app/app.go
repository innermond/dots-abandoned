package app

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/enc/branca"
	"github.com/innermond/dots/env"
)

var err error

// dependencies initialized once when app is started
// private variables that exists for the lifetime of application
var db *sql.DB

func database() error {
	// mysql database
	db, err = sql.Open("mysql", env.Dsn())
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func Db() *sql.DB {
	return db
}

var tok enc.Tokenizer

func tokenizer() {
	tok = branca.NewEncrypt(env.TokenKey(), time.Second*10)
}

func Tok() enc.Tokenizer {
	return tok
}

func run() {
	err = database()
	if err != nil {
		log.Fatal(err)
	}
	tokenizer()
}

var once sync.Once

func Init() {
	once.Do(run)
}
