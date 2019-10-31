package app

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"

	"github.com/almerlucke/go-iban/iban"
	"github.com/innermond/dots"
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

func validCompanyRegister(cd InputCompanyRegister) error {
	var err error
	if err = validCompany(cd.Company); err != nil {
		return err
	}

	for _, a := range cd.Addresses {
		if err = validAddress(a); err != nil {
			return err
		}
	}

	for _, b := range cd.Ibans {
		if err = validIban(b); err != nil {
			return err
		}
	}

	return nil
}

func validCompany(c dots.Company) error {
	var err error
	err = between("longname", c.Longname, 4, 50)
	if err != nil {
		return err
	}
	err = between("tin", c.TIN, 4, 30)
	if err != nil {
		return err
	}
	err = between("rn", c.RN, 4, 30)
	if err != nil {
		return err
	}
	// printable chars
	ss := []string{c.Longname, c.TIN, c.RN}
	for _, s := range ss {
		if !isPrintable(s) {
			return fmt.Errorf("%s contains invalid characters", s)
		}
	}
	return nil
}

func validAddress(a dots.Address) error {
	return nil
}

func validIban(b dots.Iban) error {
	_, err := iban.NewIBAN(b.Iban)
	if err != nil {
		return err
	}
	return nil
}

func between(name, s string, a, z int) error {
	lens := len(s)
	if false == (a <= lens && lens <= z) {
		return fmt.Errorf("%s length %d between %d and %d", name, lens, a, z)
	}
	return nil
}

func isPrintable(s string) bool {
	for _, c := range s {
		if !unicode.IsPrint(c) {
			return false
		}
	}
	return true
}
