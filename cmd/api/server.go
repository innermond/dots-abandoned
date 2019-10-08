package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/innermond/dots"
	"github.com/innermond/dots/service"
)

type server struct {
	*http.Server

	db *sql.DB
}

func (s *server) checkHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		out := output{Payload: serverHealth, Code: http.StatusOK}
		out.into(r, w)
	}
}

func (s *server) userPost() http.HandlerFunc {

	type response struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	userService := service.User(s.db)

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// out is the response
		var out output
		defer out.into(r, w)

		// input data to app data (json to struct)
		var ud dots.User
		if err := json.NewDecoder(r.Body).Decode(&ud); err != nil {
			out = output{Payload: err, Code: http.StatusBadRequest}
			return
		}

		// send app data to service layer
		newid, err := userService.Add(ud)
		if err != nil {
			out = output{err, http.StatusInternalServerError}
			return
		}

		// response
		resp := response{
			ID:       newid,
			Username: ud.Username,
		}

		out = output{resp, http.StatusOK}
	}

}

func (s *server) login() http.HandlerFunc {

	type response struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	userService := service.User(s.db)

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// out is the response
		var out output
		defer out.into(r, w)

		// input data to app data (json to struct)
		ud := &dots.User{}
		err := error(nil)

		if err = json.NewDecoder(r.Body).Decode(ud); err != nil {
			out = output{Payload: err, Code: http.StatusBadRequest}
			return
		}

		// send app data to service layer
		ud, err = userService.FindByUsername(ud.Username)
		if err != nil {
			out = output{err, http.StatusInternalServerError}
			return
		}

		// response
		resp := response{
			ID:       ud.ID,
			Username: ud.Username,
		}

		out = output{resp, http.StatusOK}
	}

}
