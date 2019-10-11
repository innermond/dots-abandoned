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
