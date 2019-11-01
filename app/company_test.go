package app

import (
	"testing"

	"github.com/innermond/dots/store"
	"github.com/innermond/dots/testdata"
)

func Test_AddCompany(t *testing.T) {

	for _, tc := range testdata.CompanyRegisterValid {
		id, err := AddCompany(tc.Company)
		if err != nil {
			t.Fatal(err)
		}

		op := store.CompanyOp()
		// assure test user is deleted as at this point is surely created
		defer func(name string) {
			t.Logf("defer delete test company %s", name)
			op.Delete(id)
		}(tc.Company.Longname)

	}
}

func Test_RegisterCompany(t *testing.T) {

	for _, tc := range testdata.CompanyRegisterValid {
		id, err := RegisterCompany(InputCompanyRegister(tc))
		if err != nil {
			t.Fatal(err)
		}

		op := store.CompanyOp()
		// assure test user is deleted as at this point is surely created
		defer func(name string) {
			t.Logf("defer delete test company %s", name)
			op.Delete(id)
		}(tc.Company.Longname)

	}
}
