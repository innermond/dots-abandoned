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

type this struct {
	db  *sql.DB
	tok enc.Tokenizer
}

var app this

var err error

// dependencies initialized once when app is started
// private variables that exists for the lifetime of application

func database() error {
	// mysql database
	app.db, err = sql.Open("mysql", env.Dsn())
	if err != nil {
		return err
	}
	if err = app.db.Ping(); err != nil {
		return err
	}
	return nil
}

func Db() *sql.DB {
	return app.db
}

func tokenizer() {
	app.tok = branca.NewEncrypt(env.TokenKey(), time.Second*10)
}

func Tok() enc.Tokenizer {
	return app.tok
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
