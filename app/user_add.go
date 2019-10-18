package app

import (
	"errors"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/store"
)

func AddUser(u dots.User) (int, error) {
	var err error
	// store an encrypted password
	u.Password, err = enc.Password(u.Password)
	if err != nil {
		return 0, err
	}

	// validate
	lenu := len(u.Username)
	if lenu > 16 || lenu < 8 {
		return 0, errors.New("app.AddUser: invalid username")
	}
	return store.UserOp().Add(u)
}
