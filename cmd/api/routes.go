package main

import (
	"log"

	"github.com/go-chi/chi"
)

func (s *server) routes() {
	router, ok := s.Handler.(*chi.Mux)
	if !ok {
		log.Fatal("no expected router")
	}
	router.Route(API_PATH, func(x chi.Router) {
		x.Post("/login", jzon(s.login()))
		x.Post("/user", jzon(s.userPost()))
		x.Get("/health", jzon(s.checkHealth()))
	})
}
