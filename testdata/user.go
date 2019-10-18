package testdata

import "github.com/innermond/dots/rand"

var UserPassword = []struct {
	Usr string
	Pwd string
}{
	{rand.Letters(6), rand.Letters(10)},
	{rand.Letters(8), rand.Letters(20)},
	{rand.Letters(8), rand.Letters(30)},
}
