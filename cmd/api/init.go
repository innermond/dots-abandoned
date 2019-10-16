package main

import (
	"github.com/innermond/dots/enc"
	"github.com/innermond/dots/env"
	"github.com/innermond/dots/store"
)

func init() {
	env.Init()
	store.Init()
	enc.Init()
}
