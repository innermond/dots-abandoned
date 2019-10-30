package app

import (
	"testing"

	"github.com/innermond/dots/store"
	"github.com/innermond/dots/testdata"
)

func Test_AddCompany(t *testing.T) {

	for _, tc := range testdata.Company {
		id, err := AddCompany(tc)
		if err != nil {
			t.Fatal(err)
		}

		op := store.CompanyOp()
		// assure test user is deleted as at this point is surely created
		defer func(name string) {
			t.Logf("defer delete test company %s", name)
			op.Delete(id)
		}(tc.Longname)

	}
}
