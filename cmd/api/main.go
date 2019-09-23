package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var serverHealth int32

func main() {
	// mysql database
	dsn := "root:M0b1d1c3@/printoo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	s := &server{
		Server: &http.Server{
			Addr:              ":2000",
			Handler:           r,
			ReadHeaderTimeout: time.Second * 5,
			IdleTimeout:       time.Second * 30,
		},
		db: db,
	}

	s.routes()

	s.RegisterOnShutdown(func() {
		log.Println("Server is cold")
	})

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
	log.Printf("server started on %s\n", s.Addr)
	atomic.StoreInt32(&serverHealth, 1)
	// blocks & servs
	if err := s.ListenAndServeTLS("./server.crt", "./server.key"); err != http.ErrServerClosed {
		log.Fatalf("server cannot start %v\n", err)
	}

	// quiting blocks here until server gracefully has closed
	<-done
	log.Println("Server stopped")

}
