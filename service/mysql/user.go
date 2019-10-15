package mysql

import (
	"database/sql"
	"strconv"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
)

// implements dots.userService interface
type userService struct {
	db *sql.DB
}

func UserService(db *sql.DB) dots.UserService {
	return &userService{db}
}

type UserStore = userService

func User(db *sql.DB) *UserStore {
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

func (s *userService) Register(u dots.User, role dots.Role) (int, error) {
	qryUser := "insert into users (username, password) values(?, ? )"
	qryUserRole := "insert into user_roles (user_id, role_name) values(?, ? )"
	db := s.db

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmUser, err := tx.Prepare(qryUser)
	if err != nil {
		return 0, err
	}
	defer stmUser.Close()

	resUser, err := stmUser.Exec(u.Username, u.Password)
	if err != nil {
		return 0, err
	}
	uid, err := resUser.LastInsertId()

	stmUserRole, err := tx.Prepare(qryUserRole)
	if err != nil {
		return 0, err
	}
	defer stmUserRole.Close()

	_, err = stmUserRole.Exec(uid, role)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(uid), err
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

// Login logins a dot.User
func (s *userService) Login(uname, pwd string, tokenizer enc.Tokenizer) (token string, err error) {

	var u = new(dots.User)

	u, err = s.FindByUsername(uname)
	if err != nil {
		return
	}

	err = enc.HashIsPassword(u.Password, pwd)
	if err != nil {
		return
	}

	token, err = tokenizer.Encode(strconv.Itoa(u.ID))
	if err != nil {
		return
	}

	return
}
