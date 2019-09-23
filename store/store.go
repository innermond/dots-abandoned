package store

import (
	"github.com/innermond/dots"
)

type User interface {
	Add(dots.User) (int64, error)
}
