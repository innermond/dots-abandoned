package app

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/store"
)

func AddUser(u dots.User) (int, error) {
	var err error
	// validate
	// password length
	lenp := len(u.Password)
	if false == (8 <= lenp && lenp <= 16) {
		return 0, fmt.Errorf("app.AddUser: password length %d between 8 and 16", lenp)
	}
	var rx = "^.*[A-Z].*" // at least one uppercase character
	re := regexp.MustCompile(rx)
	if !re.MatchString(u.Password) {
		return 0, errors.New("must contain at least one uppercase character")
	}
	rx = ".*[!@#$&].*" // at least one special character
	re = regexp.MustCompile(rx)
	if !re.MatchString(u.Password) {
		return 0, errors.New("must contain at least one special character")
	}
	rx = ".*[0-9].*" // at least one digit
	re = regexp.MustCompile(rx)
	if !re.MatchString(u.Password) {
		return 0, errors.New("must contain at least one digit character")
	}

	// store an encrypted password
	u.Password, err = enc.Password(u.Password)
	if err != nil {
		return 0, err
	}

	// username length
	lenu := len(u.Username)
	if !(8 <= lenu && lenu <= 16) {
		return 0, errors.New("app.AddUser: username length between 8 and 16")
	}
	// username contains only letters and numbers and underscore
	for _, c := range u.Username {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c) || c == rune('_')
		if !ok {
			return 0, errors.New("allowed only letters, numbers and underscores")
		}
	}
	// username quality
	uniqq := make(map[rune]bool)
	numUniqq := 0
	for _, c := range u.Username {
		if _, found := uniqq[c]; !found {
			uniqq[c] = true
			numUniqq++
		}
	}
	if numUniqq < 4 {
		return 0, errors.New("too repetitiv; at least 4 characters mus be uniques")
	}

	return store.UserOp().Add(u)
}
