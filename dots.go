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
