package store

import (
	"testing"

	"github.com/innermond/dots/testdata"
)

func Test_IbansAdd(t *testing.T) {

	op := CompanyOp()

	company := testdata.CompanyRegisterValid[0].Company
	cid, err := op.Add(company)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testdata.CompanyRegisterValid {
		ids, err := op.IbansAdd(cid, tc.Ibans)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("delete test ibans %v", ids)
		op.IbansDelete(cid, ids)
	}

	// delete company will trigger addresses delete
	// placed here in order to test AddressesDelete
	defer func(companyName string) {
		t.Logf("defer delete test company %s %d", companyName, cid)
		op.Delete(cid)
	}(company.Longname)
}

func Test_IbanModify(t *testing.T) {

	op := CompanyOp()

	company := testdata.CompanyRegisterValid[0].Company
	cid, err := op.Add(company)
	if err != nil {
		t.Fatal(err)
	}

	defer func(companyName string) {
		t.Logf("defer delete test company %s %d", companyName, cid)
		op.Delete(cid)
	}(company.Longname)

	func() {
		for _, tc := range testdata.CompanyRegisterValid {
			ids, err := op.IbansAdd(cid, tc.Ibans)
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				t.Logf("delete test ibans %v", ids)
				op.AddressesDelete(cid, ids)
			}()

			// assume range returns same order as ids
			for i, iban := range tc.Ibans {
				iban.ID = ids[i]
				err = op.IbanModify(cid, iban)
				if err != nil {
					t.Fatal(err)
				}
			}
		}
	}()
}
