package store

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots/testdata"
)

func Test_CompanyRegister(t *testing.T) {
	op := CompanyOp()

	for _, tc := range testdata.CompanyRegisterValid {
		id, err := op.Register(tc.Company, tc.Addresses, tc.Ibans)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(companyName string) {
			t.Logf("defer delete test company %s", companyName)
			op.Delete(id)
		}(tc.Company.Longname)
	}
}

func Test_CompanyModify(t *testing.T) {
	op := CompanyOp()

	for _, tc := range testdata.CompanyRegisterValid {
		id, err := op.Add(tc.Company)
		if err != nil {
			t.Fatal(err)
		}
		tc.Company.ID = id
		tc.Company.Longname += "modified"
		tc.Company.IsClient = true

		err = op.Modify(tc.Company)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(companyName string) {
			t.Logf("defer delete test company %s", companyName)
			op.Delete(id)
		}(tc.Company.Longname)
	}
}
