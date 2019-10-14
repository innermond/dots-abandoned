package main

import (
	"github.com/innermond/dots/app"
	"github.com/innermond/dots/env"
)

func init() {
	env.Init()
	app.Init()
}
