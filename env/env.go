package env

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const API_PATH = "/api/v1"

func Get(key, alternative string) (string, error) {
	err = fmt.Errorf("%s not found", key)
	val, found := os.LookupEnv(key)
	if !found {
		if alternative != "" {
			return alternative, nil
		}
		return "", err
	}
	return val, nil
}

var (
	err                       error
	debug                     bool
	host, port, dsn, tokenKey string
)

func flagParse() error {
	dsn, err = Get("DOTS_DSN", "")
	if err != nil || dsn == "" {
		return errors.New("dsn not received")
	}
	flag.BoolVar(&debug, "debug", false, "activate debug")
	flag.StringVar(&host, "h", "", "host address")
	flag.StringVar(&port, "p", "2000", "port part of server's address")
	flag.StringVar(&dsn, "dsn", dsn, "database DSN string")
	flag.StringVar(&tokenKey, "tokenKey", strings.Repeat("x", 32), "key used tokens encryption")
	flag.Parse()

	host = strings.TrimRight(host, ":")
	if host != "" {
		host, port, err = net.SplitHostPort(host)
		if err != nil {
			return err
		}
	}
	return nil
}

var db *sql.DB

func database() error {
	// mysql database
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func init() {
	log.Println("env.init start")
	err = flagParse()
	if err != nil {
		log.Fatal(err)
	}
	err = database()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("env.init done")
}

func Db() *sql.DB {
	return db
}

func Addr() string {
	return host + ":" + port
}

func Debug() bool {
	return debug
}

func TokenKey() string {
	return tokenKey
}
