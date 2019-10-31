package app

import (
	"github.com/innermond/dots"
	"github.com/innermond/dots/store"
)

func AddCompany(c dots.Company) (int, error) {
	return store.CompanyOp().Add(c)
}

type InputCompanyRegister struct {
	Company   dots.Company
	Addresses []dots.Address
	Ibans     []dots.Iban
}

func RegisterCompany(cd InputCompanyRegister) (int, error) {
	// validation
	err := validCompanyRegister(cd)
	if err != nil {
		return 0, err
	}

	c := cd.Company
	aa := cd.Addresses
	bb := cd.Ibans

	return store.CompanyOp().Register(c, aa, bb)
}
