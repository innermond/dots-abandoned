package testdata

type UserData struct {
	Usr string
	Pwd string
}

var UserPassword = []UserData{
	{"gbbg1_2434", "a!sa_Ar3tQ"},
	{"ghgdh23_34", "%ZEP_a!sOP1"},
}

var UserInvalid = []UserData{
	{"short", "qqqq"},
}
