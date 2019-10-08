package service

import (
	"database/sql"

	"github.com/innermond/dots"
	store "github.com/innermond/dots/service/mysql"
)

func User(db *sql.DB) dots.UserService {
	return store.UserService(db)
}
