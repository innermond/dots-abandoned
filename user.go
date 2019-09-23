package dots

type UserService interface {
	Add(User) error
}

type User struct {
	ID       int
	Username string
	Password string
}
