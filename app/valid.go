package app

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

func validUser(u InputUserRegister) error {
	// validate
	// password length
	lenp := len(u.Password)
	if false == (8 <= lenp && lenp <= 16) {
		return fmt.Errorf("app.AddUser: password length %d between 8 and 16", lenp)
	}
	var rx = "^.*[A-Z].*" // at least one uppercase character
	re := regexp.MustCompile(rx)
	if !re.MatchString(u.Password) {
		return errors.New("must contain at least one uppercase character")
	}
	rx = ".*[!@#$&].*" // at least one special character
	re = regexp.MustCompile(rx)
	if !re.MatchString(u.Password) {
		return errors.New("must contain at least one special character")
	}
	rx = ".*[0-9].*" // at least one digit
	re = regexp.MustCompile(rx)
	if !re.MatchString(u.Password) {
		return errors.New("must contain at least one digit character")
	}

	// username length
	lenu := len(u.Username)
	if !(8 <= lenu && lenu <= 16) {
		return errors.New("app.AddUser: username length between 8 and 16")
	}
	// username contains only letters and numbers and underscore
	for _, c := range u.Username {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c) || c == rune('_')
		if !ok {
			return errors.New("allowed only letters, numbers and underscores")
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
		return errors.New("too repetitiv; at least 4 characters mus be uniques")
	}
	return nil
}
