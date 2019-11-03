package dots

type UserService interface {
	Add(User) (int, error)
	//FindByUsername(string) (*User, error)
}

type User struct {
	ID       int
	Username string
	Password string
}

type Role string

const (
	AnonymousRole  Role = "anonymous"
	UserRole       Role = "user"
	AdminRole      Role = "admin"
	SuperAdminRole Role = "superadmin"
)

type Company struct {
	ID       int
	Longname string
	TIN      string
	RN       string
}

type Address struct {
	ID       int
	Address  string
	Location Point
}

type Point struct {
	X, Y float64
}

type Iban struct {
	ID       int
	Iban     string
	Bankname string
}
