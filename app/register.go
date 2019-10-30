package app

import (
	"strconv"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/store"
)

type InputUserRegister struct {
	Username string
	Password string
}

func Register(ud InputUserRegister) (string, error) {
	// verify UserData is valid
	err := validUser(ud)
	if err != nil {
		return "", err
	}
	// store UserData
	encrypted, err := enc.Password(ud.Password)
	if err != nil {
		return "", err
	}

	u := dots.User{
		Username: ud.Username,
		Password: encrypted,
	}
	uid, err := store.UserOp().Register(u, dots.UserRole)
	if err != nil {
		return "", err
	}

	// give a token
	uidstr := strconv.Itoa(uid)
	tok, err := enc.Tok().Encode(uidstr)
	if err != nil {
		return "", err
	}
	return tok, nil
}
