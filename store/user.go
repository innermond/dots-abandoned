package store

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/innermond/dots"
	"github.com/innermond/dots/enc"
)

type userOp struct {
}

var userOperations *userOp

func UserOp() *userOp {
	if userOperations == nil {
		userOperations = &userOp{}
	}
	return userOperations
}

func (op *userOp) Add(u dots.User) (int, error) {
	qry := "insert into users (username, password) values(?, ? )"
	db := store.DB

	res, err := db.Exec(qry, u.Username, u.Password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (op *userOp) Delete(uid int) error {
	qry := "delete from users where id = ? limit 1"
	db := store.DB
	err := error(nil)
	_, err = db.Exec(qry, uid)
	if err != nil {
		return err
	}
	return nil
}

func (op *userOp) Unrole(uid int, role []dots.Role) error {
	var qry string
	lenrole := len(role)
	if lenrole == 1 {
		qry = "delete from user_roles where user_id = ? and role_name = ? limit 1"
	} else {
		in := strings.Repeat(",?", lenrole)[1:]
		qry = fmt.Sprintf("delete from user_roles where user_id = ? and role_name in(%s)", in)
	}
	db := store.DB
	err := error(nil)
	var args []interface{}
	args = append(args, uid)
	for _, rol := range role {
		args = append(args, interface{}(rol))
	}
	_, err = db.Exec(qry, args...)
	if err != nil {
		return err
	}
	return nil
}

func (op *userOp) Register(u dots.User, role dots.Role) (int, error) {
	qryUser := "insert into users (username, password) values(?, ? )"
	qryUserRole := "insert into user_roles (user_id, role_name) values(?, ? )"
	db := store.DB

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	resUser, err := tx.Exec(qryUser, u.Username, u.Password)
	if err != nil {
		return 0, err
	}
	uid, err := resUser.LastInsertId()

	_, err = tx.Exec(qryUserRole, uid, role)
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
func (op *userOp) FindAll() ([]dots.User, error) {
	qry := "select id, username, password from users"
	db := store.DB

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
func (op *userOp) FindByUsername(uname string) (*dots.User, error) {
	qry := "select id, username, password from users where username = ?  limit 1"
	db := store.DB

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
func (op *userOp) Login(uname, pwd string, tokenizer enc.Tokenizer) (token string, err error) {

	var u = new(dots.User)

	u, err = op.FindByUsername(uname)
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
