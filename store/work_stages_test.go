package store

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots"
	"github.com/innermond/dots/testdata"
)

func Test_WorkStagesAdd(t *testing.T) {
	op := WorkStageOp()

	for _, tc := range testdata.WorkStagesValid {
		id, err := op.Add(tc)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("added stage %s with id %d", tc.Stage, id)

		// assure test user is deleted as at this point is surely created
		defer func(tc dots.WorkStage) {
			t.Logf("defer delete test workStage %s", tc.Stage)
			op.Delete(tc.Ordered)
		}(tc)

		w, err := op.FindByOrdered(tc.Ordered)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("stage %v", w.Stage)
	}
}

/*
func Test_WorkStagesModify(t *testing.T) {
	op := WorkOp()

	for _, tc := range testdata.WorkStagesValid {
		id, err := op.Add(tc)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v", tc)
		tc.ID = id
		tc.Label += "modified"
		tc.UnitPrice.Add(&tc.UnitPrice)
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
*/
