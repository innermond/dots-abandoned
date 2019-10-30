package app

import (
	"github.com/innermond/dots"
	"github.com/innermond/dots/store"
)

func AddCompany(c dots.Company) (int, error) {
	return store.CompanyOp().Add(c)
}
