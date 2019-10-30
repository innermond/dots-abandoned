package testdata

import "github.com/innermond/dots"

type UserData struct {
	Usr  string
	Pwd  string
	Role dots.Role
}

var UserPassword = []UserData{
	{"gbbg1_2434", "a!sa_Ar3tQ", dots.UserRole},
	{"ghgdh23_34", "%ZEP_a!sOP1", dots.AdminRole},
}

var UserInvalid = []UserData{
	{"short", "qqqq", dots.AnonymousRole},
}
