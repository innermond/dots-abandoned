package env

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	err                       error
	debug                     bool
	host, port, dsn, tokenKey string
)

var (
	ErrDsnEmpty = errors.New("env: dsn not set")
)

func flagParse() error {
	dsn, err = Get("DOTS_DSN", "")
	if err != nil {
		return err
	}
	if dsn == "" {
		return ErrDsnEmpty
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

func Addr() string {
	return host + ":" + port
}

func Debug() bool {
	return debug
}

func TokenKey() string {
	return tokenKey
}

func Dsn() string {
	return dsn
}

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

func run() {
	log.Println("env.init start")
	err = flagParse()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("env.init done")
}

var once sync.Once

func Init() {
	once.Do(run)
}
