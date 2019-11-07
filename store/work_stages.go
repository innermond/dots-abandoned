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

	res, err := db.Exec(qry, ws.Stage, ws.Description, ws.Ordered)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func (op *workStageOp) Delete(id int) error {
	qry := "delete from work_stages where id = ? limit 1"
	db := store.DB
	err := error(nil)
	_, err = db.Exec(qry, id)
	if err != nil {
		return err
	}
	return nil
}

func (op workStageOp) FindById(id int) (dots.WorkStage, error) {
	qry := "select stage, description, ordered from work_stages where id= ? limit 1"
	db := store.DB

	var ws = dots.WorkStage{}

	err := db.QueryRow(qry, id).Scan(&ws.Stage, &ws.Description, &ws.Ordered)
	if err == sql.ErrNoRows {
		return ws, err
	}
	if err != nil {
		return ws, err
	}

	ws.ID = id
	return ws, nil
}

func (op *workStageOp) Modify(ws dots.WorkStage) error {
	qry := "update work_stages set stage=?, description=?, ordered=? where id=? limit 1"
	db := store.DB

	_, err := db.Exec(qry, ws.Stage, ws.Description, ws.Ordered, ws.ID)
	if err != nil {
		return err
	}

	return nil
}
