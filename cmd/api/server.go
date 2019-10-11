package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/service"
	store "github.com/innermond/dots/service/mysql"
)

type server struct {
	*http.Server

	db        *sql.DB
	tokenizer enc.Tokenizer
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
		var ud inputUserAutomodifying
		if err := json.NewDecoder(r.Body).Decode(&ud); err != nil {
			out = output{Payload: err, Code: http.StatusBadRequest}
			return
		}

		// send app data to service layer
		newid, err := userService.Add(dots.User(ud))
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

	type token struct {
		Token string `json:"token"`
	}

	userStore := store.User(s.db)

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
		var tk string
		tk, err = userStore.Login(ud.Username, ud.Password, s.tokenizer)
		if err != nil {
			out = output{err, http.StatusInternalServerError}
			return
		}

		// response
		resp := token{Token: tk}

		out = output{resp, http.StatusOK}
	}

}

//TODO password encrypting belongs to service layer
type inputUserAutomodifying dots.User

func (u *inputUserAutomodifying) UnmarshalJSON(data []byte) error {
	var err error

	type input inputUserAutomodifying
	out := (*input)(u)

	if err = json.Unmarshal(data, out); err != nil {
		return err
	}
	out.Password, err = enc.Password(u.Password)
	if err != nil {
		return err
	}
	return nil
}
