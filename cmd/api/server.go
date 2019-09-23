package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/innermond/dots"
	"github.com/innermond/dots/service"
	"github.com/innermond/dots/store/mysql"
)

type server struct {
	*http.Server

	db *sql.DB
}

func (s *server) handleHealthGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		out := output{Payload: serverHealth, Code: http.StatusOK}
		out.into(r, w)
	}
}

func (s *server) handleUserPost() http.HandlerFunc {

	type response struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	userStore := mysql.NewUser(s.db)
	userService := service.NewUser(userStore)

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var out output

		var ud dots.User
		if err := json.NewDecoder(r.Body).Decode(&ud); err != nil {
			out = output{Payload: err, Code: http.StatusBadRequest}
			out.into(r, w)
			return
		}

		newid, err := userService.Add(ud)
		if err != nil {
			out = output{err, http.StatusInternalServerError}
			out.into(r, w)
			return
		}

		resp := response{
			ID:       newid,
			Username: ud.Username,
		}

		out = output{resp, http.StatusOK}
		out.into(r, w)
	}

}
