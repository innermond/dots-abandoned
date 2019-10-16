package app

import (
	"strconv"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/store"
)

// Login logins a dot.User
// db and tok are private vars to app package
// initialized when app started
func Login(uname, pwd string) (token string, err error) {

	var u = new(dots.User)

	u, err = store.UserOp().FindByUsername(uname)
	if err != nil {
		return
	}

	err = enc.HashIsPassword(u.Password, pwd)
	if err != nil {
		return
	}

	token, err = enc.Tok().Encode(strconv.Itoa(u.ID))
	if err != nil {
		return
	}

	return
}
