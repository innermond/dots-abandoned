package app

import (
	"strconv"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	store "github.com/innermond/dots/service/mysql"
)

type InputUserRegister struct {
	Username string
	Password string
}

// TODO when register store role
func Register(ud InputUserRegister) (string, error) {
	// verify UserData is valid

	// store UserData
	encrypted, err := enc.Password(ud.Password)
	if err != nil {
		return "", err
	}

	u := dots.User{
		Username: ud.Username,
		Password: encrypted,
	}
	uid, err := store.User(app.db).Register(u, dots.UserRole)
	if err != nil {
		return "", err
	}

	// give a token
	uidstr := strconv.Itoa(uid)
	tok, err := app.tok.Encode(uidstr)
	if err != nil {
		return "", err
	}
	return tok, nil
}
