package store

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots"
	"github.com/innermond/dots/env"
	"github.com/innermond/dots/testdata"
)

func TestMain(m *testing.M) {
	env.Init()
	Init()
	os.Exit(m.Run())
	defer Close()
}

func Test_UserAdd(t *testing.T) {
	op := UserOp()

	for _, tc := range testdata.UserPassword {
		ua := dots.User{Username: tc.Usr, Password: tc.Pwd}
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

func Test_UserRegister(t *testing.T) {
	op := UserOp()

	for _, tc := range testdata.UserPassword {
		ua := dots.User{Username: tc.Usr, Password: tc.Pwd}
		id, err := op.Register(ua, tc.Role)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(usr string) {
			t.Logf("defer delete test user %s", usr)
			op.Delete(id)
			op.Unrole(id, []dots.Role{tc.Role, dots.UserRole})
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
