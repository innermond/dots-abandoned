package store

import (
	"math/rand"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/env"
)

func init() {
	env.Init()
	Init()
	enc.Init()
}

func Test_AddUser(t *testing.T) {
	tt := []struct {
		usr string
		pwd string
	}{
		{letters(6), letters(10)},
		{letters(8), letters(20)},
		{letters(8), letters(40)},
	}
	db := store.DB
	defer db.Close()
	err := error(nil)

	for _, tc := range tt {
		ua := dots.User{Username: tc.usr, Password: tc.pwd}
		plain := ua.Password
		ua.Password, err = enc.Password(ua.Password)
		if err != nil {
			t.Fatal(err)
		}
		id, err := UserOp().Add(ua)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(usr string) {
			t.Logf("defer delete test user %s", usr)
			_, err = db.Exec("delete from users where id = ? limit 1", id)
			if err != nil {
				t.Fatal(err)
			}
		}(tc.usr)

		uz, err := UserOp().FindByUsername(tc.usr)
		if err != nil {
			t.Fatal(err)
		}
		err = enc.HashIsPassword(uz.Password, plain)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("hash %s is password %s", uz.Password, plain)
	}
}

func Test_LoginUser(t *testing.T) {
	//tokenKey := strings.Repeat("x", 32)
	//tokenizer := branca.NewEncrypt(tokenKey, time.Second*10)
	t.Skip()
}

// letters printable ascii string; lower letters
func letters(n int) string {
	x := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]byte, n)
	b := 65
	for i := 0; i < n; i++ {
		// only from 65-122 inclusiv; letters lower and upper case
		// except range 90-97 exclusive; printable chars but not letters
		b = x.Intn(123-65) + 65
		for b > 90 && b < 97 {
			b = x.Intn(123-65) + 65
		}
		out = append(out, byte(b))
	}
	return string(out)
}

func Benchmark_letters(b *testing.B) {
	bench_letters(8, b)
	bench_letters(6, b)
}

var result_letters string

func bench_letters(i int, b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = letters(i)
	}
	result_letters = r
}
