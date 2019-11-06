package store

import (
	"math/big"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots"
	"github.com/innermond/dots/testdata"
)

func Test_WorkAdd(t *testing.T) {
	op := WorkOp()

	for _, tc := range testdata.WorkValid {
		t.Logf("unitprice %v", tc.UnitPrice)
		id, err := op.Add(tc)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(label string) {
			t.Logf("defer delete test work %s", label)
			op.Delete(id)
		}(tc.Label)

		w, err := op.FindById(id)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("unitprice %v", w.UnitPrice)
	}
}

func Test_WorkModify(t *testing.T) {
	op := WorkOp()

	for _, tc := range testdata.WorkValid {
		id, err := op.Add(tc)
		if err != nil {
			t.Fatal(err)
		}
		tc.ID = id
		tc.Label += "modified"
		br := big.Rat(tc.UnitPrice)
		br.Add(&br, big.NewRat(2000, 100))
		tc.UnitPrice = dots.Rational(br)

		err = op.Modify(tc)
		if err != nil {
			t.Fatal(err)
		}

		defer func(tc dots.Work) {
			t.Logf("defer delete test work %s unit price %v", tc.Label, tc.UnitPrice)
			op.Delete(id)
		}(tc)
	}
}
