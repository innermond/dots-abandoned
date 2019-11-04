package store

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots/testdata"
)

func Test_Work(t *testing.T) {
	op := WorkOp()

	for _, tc := range testdata.WorkValid {
		id, err := op.Add(tc)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(label string) {
			t.Logf("defer delete test work %s", label)
			op.Delete(id)
		}(tc.Label)
	}
}

/*
func Test_WorkModify(t *testing.T) {
	op := WorkOp()

	for _, tc := range testdata.WorkValid {
		id, err := op.Add(tc.Work)
		if err != nil {
			t.Fatal(err)
		}
		tc.Work.ID = id
		tc.Work.Longname += "modified"
		tc.Work.IsClient = true

		err = op.Modify(tc.Work)
		if err != nil {
			t.Fatal(err)
		}

		// assure test user is deleted as at this point is surely created
		defer func(companyName string) {
			t.Logf("defer delete test company %s", companyName)
			op.Delete(id)
		}(tc.Work.Longname)
	}
}
*/
