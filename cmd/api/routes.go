package main

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const API_PATH = "/api/v1"

func (s *server) routes() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	auth := s.guard()

	r.Route(API_PATH, func(x chi.Router) {
		x.Use(jzon)
		x.Post("/login", s.login())
		x.Post("/register", s.register())
		x.With(auth).Post("/user", s.userPost())
		x.Post("/company", s.companyRegister())
		x.Get("/health", s.checkHealth())
	})

	s.Server.Handler = r
}
