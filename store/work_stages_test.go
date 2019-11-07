package store

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/innermond/dots"
	"github.com/innermond/dots/testdata"
)

func addWorkStage(op *workStageOp, tc dots.WorkStage, t *testing.T) (int, func()) {
	id, err := op.Add(tc)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("added stage %s with order %d", tc.Stage, id)

	return id, func() {
		t.Logf("delete %d", id)
		op.Delete(id)
	}
}

func findWorkStage(op *workStageOp, id int, t *testing.T) {
	w, err := op.FindById(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("found stage %v %d", w.Stage, w.ID)
}

func Test_WorkStagesAdd(t *testing.T) {
	op := WorkStageOp()

	for _, tc := range testdata.WorkStagesValid {
		tid, xfunc := addWorkStage(op, tc, t)
		defer xfunc()
		findWorkStage(op, tid, t)
	}
}

func Test_WorkStagesModify(t *testing.T) {
	op := WorkStageOp()

	var err error
	for _, tc := range testdata.WorkStagesValid {
		tid, xfunc := addWorkStage(op, tc, t)

		tc.Stage += "--modified"
		tc.Ordered -= 1000000
		t.Logf("%+v", tid)
		xfunc()

		err = op.Modify(tc)
		if err != nil {
			t.Fatal(err)
		}
	}
}
