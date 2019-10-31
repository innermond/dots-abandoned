package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/innermond/dots"
	"github.com/innermond/dots/app"
	"github.com/innermond/dots/enc"
)

type server struct {
	*http.Server

	tokenizer enc.Tokenizer
}

func (s *server) checkHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		out(serverHealth, http.StatusOK).into(r, w)
	}
}

func (s *server) userPost() http.HandlerFunc {

	type response struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

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
		newid, err := app.AddUser(ud)
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
		Username string `json:"username"`
		Token    string `json:"token"`
	}

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
		tok, err := app.Login(ud.Username, ud.Password)
		if err != nil {
			out = output{err, http.StatusUnauthorized}
			return
		}

		// response
		resp := response{
			Username: ud.Username,
			Token:    tok,
		}

		out = output{resp, http.StatusOK}
	}

}

func (s *server) register() http.HandlerFunc {

	type response struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// out is the response
		var out output
		defer out.into(r, w)

		// input data to app data (json to struct)
		ud := app.InputUserRegister{}
		err := error(nil)
		tk := ""

		if err = json.NewDecoder(r.Body).Decode(&ud); err != nil {
			out = output{Payload: err, Code: http.StatusBadRequest}
			return
		}

		// send app data to service layer
		tk, err = app.Register(ud)
		if err != nil {
			out = output{err, http.StatusInternalServerError}
			return
		}

		// response
		resp := response{Username: ud.Username, Token: tk}
		out = output{resp, http.StatusOK}
	}

}

func (s *server) companyRegister() http.HandlerFunc {

	type response struct {
		CompanyID int `json:"company_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// out is the response
		var out output
		defer out.into(r, w)

		// input data to app data (json to struct)
		cd := app.InputCompanyRegister{}
		err := error(nil)

		if err = json.NewDecoder(r.Body).Decode(&cd); err != nil {
			out = output{Payload: err, Code: http.StatusBadRequest}
			return
		}

		// send app data to service layer
		id, err := app.RegisterCompany(cd)
		if err != nil {
			out = output{err.Error(), http.StatusInternalServerError}
			return
		}

		// response
		resp := response{CompanyID: id}
		out = output{resp, http.StatusOK}
	}
}

type (
	nextler func(http.Handler) http.Handler
)

var ctxKeyToken = ctxkey("token")

func (s *server) guard() nextler {

	var errEmptyToken = errors.New("empty token")

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			tok, err := decodeToken(r, s.tokenizer)
			switch {
			case
				// token empty
				tok == "" && err == nil:
				out(errEmptyToken.Error(), http.StatusUnauthorized).into(r, w)
				return
			case err != nil:
				out(err, http.StatusBadRequest).into(r, w)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, ctxKeyToken, tok)

			next.ServeHTTP(w, r)
		}
		return http.Handler(http.HandlerFunc(fn))
	}
}

// decodeToken verify if token is ok
func decodeToken(r *http.Request, tokenizer enc.Tokenizer) (string, error) {
	enctok, err := authorization(r)
	// there is no token case
	if err == nil && enctok == "" {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return tokenizer.Decode(enctok)
}

// authorization gets Authorization: Bearer <value>
func authorization(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	// no token
	if ah == "" {
		return "", nil
	}
	// parts
	pp := strings.SplitN(ah, " ", 2)
	if len(pp) != 2 || strings.ToLower(pp[0]) != "bearer" {
		return "", errors.New("malformed authorization header")
	}

	return pp[1], nil
}
