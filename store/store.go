package store

import (
	"database/sql"
	"log"
	"sync"

	"github.com/innermond/dots/env"
)

type this struct {
	*sql.DB
}

var store this

func Close() {
	store.DB.Close()
}

func database() error {
	// mysql database
	db, err := sql.Open("mysql", env.Dsn())
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	store = this{db}
	return nil
}

func run() {
	err := database()
	if err != nil {
		log.Fatal(err)
	}
}

var once sync.Once

func Init() {
	once.Do(run)
}
