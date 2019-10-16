package app

import (
	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/store"
)

func AddUser(u dots.User) (int, error) {
	var err error
	u.Password, err = enc.Password(u.Password)
	if err != nil {
		return 0, err
	}
	return store.UserOp().Add(u)
}
