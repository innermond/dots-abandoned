package service

import (
	"github.com/innermond/dots"
	"github.com/innermond/dots/store"
)

type User struct {
	store store.User
}

func NewUser(store store.User) *User {
	return &User{store}
}
func (u *User) Add(ud dots.User) (int, error) {
	id, err := u.store.Add(ud)
	return int(id), err
}
