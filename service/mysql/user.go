package mysql

import (
	"database/sql"

	"github.com/innermond/dots"
)

// implements dots.userService interface
type userService struct {
	db *sql.DB
}

func UserService(db *sql.DB) dots.UserService {
	return &userService{db}
}

func (s *userService) Add(u dots.User) (int, error) {
	qry := "insert into users (username, password) values(?, ? )"
	db := s.db

	res, err := db.Exec(qry, u.Username, u.Password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

// FindAll gets all users
func (s *userService) FindAll() ([]dots.User, error) {
	qry := "select id, username, password from users"
	db := s.db

	rows, err := db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		u  = dots.User{}
		uu = []dots.User{}
	)
	for rows.Next() {
		if err = rows.Scan(&u.ID, &u.Username, &u.Password); err != nil {
			return nil, err
		}
		uu = append(uu, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return uu, nil
}

// FindByUsername return one dot.User
func (s *userService) FindByUsername(uname string) (*dots.User, error) {
	qry := "select id, username, password from users where username = ?  limit 1"
	db := s.db

	var u = new(dots.User)

	err := db.QueryRow(qry, uname).Scan(&u.ID, &u.Username, &u.Password)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
