package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/env"
	"github.com/innermond/dots/store"
)

var serverHealth int32

func main() {

	s := &server{
		Server: &http.Server{
			Addr: env.Addr(),
			//Handler:     r,
			ReadTimeout: time.Second * 10,
			//WriteTimeout:      time.Second * 10,
			ReadHeaderTimeout: time.Second * 5,
			IdleTimeout:       time.Second * 30,
		},
		tokenizer: enc.Tok(),
	}
	defer store.Close()

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
	log.Printf("server started on %s; debug mode %v\n", s.Addr, env.Debug())
	atomic.StoreInt32(&serverHealth, 1)
	// blocks & servs
	if err := s.ListenAndServeTLS("./server.crt", "./server.key"); err != http.ErrServerClosed {
		log.Fatalf("server cannot start %v\n", err)
	}

	// quiting blocks here until server gracefully has closed
	<-done
	log.Println("Server stopped")

}
