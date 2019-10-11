package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots/enc/branca"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const API_PATH = "/api/v1"

var serverHealth int32

func main() {
	var (
		debug bool

		host, port, dsn, tokenKey string

		db  *sql.DB
		err error
	)

	flag.BoolVar(&debug, "debug", false, "activate debug")
	flag.StringVar(&host, "h", "", "host address")
	flag.StringVar(&port, "p", "2000", "port part of server's address")
	flag.StringVar(&dsn, "dsn", "", "database DSN string")
	flag.StringVar(&tokenKey, "tokenKey", strings.Repeat("x", 32), "key used tokens encryption")
	flag.Parse()

	host = strings.TrimRight(host, ":")
	if host != "" {
		host, port, err = net.SplitHostPort(host)
		if err != nil {
			log.Fatal(err)
		}
	}

	dsn, err = env("DOTS_DSN", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// mysql database
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	auth := branca.NewEncrypt(tokenKey, time.Second*10)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	s := &server{
		Server: &http.Server{
			Addr:        host + ":" + port,
			Handler:     r,
			ReadTimeout: time.Second * 10,
			//WriteTimeout:      time.Second * 10,
			ReadHeaderTimeout: time.Second * 5,
			IdleTimeout:       time.Second * 30,
		},
		db:        db,
		tokenizer: auth,
	}

	s.routes()

	done := make(chan bool)

	// quiting
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		// wait here for interruption
		<-quit

		atomic.StoreInt32(&serverHealth, 0)

		// exit event occured so create context from Shutdown
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		s.SetKeepAlivesEnabled(false)
		// if ctx expires before Shutdown would do its jobs Shutdown is canceled,
		// read 'as if it have never been called'
		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("server cannot be shutdown %v/n", err)
		}
		// server is done
		close(done)
		log.Println("server going down...")
	}()

	// working
	log.Printf("server started on %s; debug mode %v\n", s.Addr, debug)
	atomic.StoreInt32(&serverHealth, 1)
	// blocks & servs
	if err := s.ListenAndServeTLS("./server.crt", "./server.key"); err != http.ErrServerClosed {
		log.Fatalf("server cannot start %v\n", err)
	}

	// quiting blocks here until server gracefully has closed
	<-done
	log.Println("Server stopped")

}

func env(key, alternative string) (string, error) {
	var err = fmt.Errorf("%s not found", key)
	val, found := os.LookupEnv(key)
	if !found {
		if alternative != "" {
			return alternative, nil
		}
		return "", err
	}
	return val, nil
}
