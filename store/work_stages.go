package store

import (
	"database/sql"

	"github.com/innermond/dots"
)

type workStageOp struct {
}

var workStageOperations *workStageOp

func WorkStageOp() *workStageOp {
	if workStageOperations == nil {
		workStageOperations = &workStageOp{}
	}
	return workStageOperations
}

func (op *workStageOp) Add(ws dots.WorkStage) (int, error) {
	qry := "insert into work_stages (stage, description, ordered) values(?, ?, ?)"
	db := store.DB

	_, err := db.Exec(qry, ws.Stage, ws.Description, ws.Ordered)
	if err != nil {
		return 0, err
	}
	return ws.Ordered, err
}

func (op *workStageOp) Delete(ordered int) error {
	qry := "delete from work_stages where ordered = ? limit 1"
	db := store.DB
	err := error(nil)
	_, err = db.Exec(qry, ordered)
	if err != nil {
		return err
	}
	return nil
}

func (op workStageOp) FindByOrdered(ordered int) (dots.WorkStage, error) {
	qry := "select stage, description, ordered from work_stages where ordered= ? limit 1"
	db := store.DB

	var ws = dots.WorkStage{}

	err := db.QueryRow(qry, ordered).Scan(&ws.Stage, &ws.Description, &ws.Ordered)
	if err == sql.ErrNoRows {
		return ws, err
	}
	if err != nil {
		return ws, err
	}

	return ws, nil
}

func (op *workStageOp) Modify(ws dots.WorkStage) error {
	qry := "update work_stages set stage=?, description=?, ordered=?"
	db := store.DB

	_, err := db.Exec(qry, ws.Stage, ws.Description, ws.Ordered)
	if err != nil {
		return err
	}

	return nil
}
