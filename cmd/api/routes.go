package main

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/innermond/dots/env"
)

func (s *server) routes() {
	router, ok := s.Handler.(*chi.Mux)
	if !ok {
		log.Fatal("no expected router")
	}

	auth := s.guard()

	router.Route(env.API_PATH, func(x chi.Router) {
		x.Use(jzon)
		x.Post("/login", s.login())
		x.With(auth).Post("/user", s.userPost())
		x.Get("/health", s.checkHealth())
	})
}
