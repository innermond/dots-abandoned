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
	router.Route("/api/v1", func(x chi.Router) {
		x.Post("/user", jzon(s.handleUserPost()))
		x.Get("/health", jzon(s.handleHealthGet()))
	})
}
