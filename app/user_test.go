package app

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/env"
	"github.com/innermond/dots/store"
	"github.com/innermond/dots/testdata"
)

func init() {
	env.Init()
	store.Init()
	enc.Init()
}

func Test_AddUser(t *testing.T) {
	defer store.Close()
	err := error(nil)

	for _, tc := range testdata.UserPassword {
		ua := dots.User{Username: tc.Usr, Password: tc.Pwd}
		if err != nil {
			t.Fatal(err)
		}
		id, err := AddUser(ua)
		if err != nil {
			t.Fatal(err)
		}

		op := store.UserOp()
		// assure test user is deleted as at this point is surely created
		defer func(usr string) {
			t.Logf("defer delete test user %s", usr)
			op.Delete(id)
		}(tc.Usr)

		uz, err := op.FindByUsername(tc.Usr)
		if err != nil {
			t.Fatal(err)
		}
		err = enc.HashIsPassword(uz.Password, ua.Password)
		if err != nil {
			t.Errorf("hash %s is not password %s", uz.Password, ua.Password)
		}
	}
}

func Test_LoginUser(t *testing.T) {
	//tokenKey := strings.Repeat("x", 32)
	//tokenizer := branca.NewEncrypt(tokenKey, time.Second*10)
	t.Skip()
}