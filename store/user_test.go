package store

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots"
	"github.com/innermond/dots/env"
	"github.com/innermond/dots/testdata"
)

func init() {
	env.Init()
	Init()
}

func Test_User(t *testing.T) {
	defer Close()
	err := error(nil)
	op := UserOp()

	for _, tc := range testdata.UserPassword {
		ua := dots.User{Username: tc.Usr, Password: tc.Pwd}
		if err != nil {
			t.Fatal(err)
		}
		id, err := op.Add(ua)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(usr string) {
			t.Logf("defer delete test user %s", usr)
			op.Delete(id)
		}(tc.Usr)

		uz, err := op.FindByUsername(tc.Usr)
		if err != nil {
			t.Fatal(err)
		}
		if uz.Password != ua.Password {
			t.Errorf("password is not hashed at store operation level: hash %s password %s", uz.Password, ua.Password)
		}
	}
}
