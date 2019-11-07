package store

import (
	"database/sql"

	"github.com/innermond/dots"
)

type workOp struct {
}

var workOperations *workOp

func WorkOp() *workOp {
	if workOperations == nil {
		workOperations = &workOp{}
	}
	return workOperations
}

func (op *workOp) Add(w dots.Work) (int, error) {
	qry := "insert into works (label, quantity, unit, unitprice, currency) values(?, ?, ?, ?, ?)"
	db := store.DB

	res, err := db.Exec(qry, w.Label, w.Quantity, w.Unit, w.UnitPrice, w.Currency)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (op *workOp) Delete(wid int) error {
	qry := "delete from works where id = ? limit 1"
	db := store.DB
	err := error(nil)
	_, err = db.Exec(qry, wid)
	if err != nil {
		return err
	}
	return nil
}

func (op workOp) FindById(wid int) (dots.Work, error) {
	qry := "select label, quantity, unit, unitprice, currency from works where id= ? limit 1"
	db := store.DB

	var w = dots.Work{}

	err := db.QueryRow(qry, wid).Scan(&w.Label, &w.Quantity, &w.Unit, &w.UnitPrice, &w.Currency)
	if err == sql.ErrNoRows {
		return w, err
	}
	if err != nil {
		return w, err
	}

	return w, nil
}

func (op *workOp) Modify(w dots.Work) error {
	qry := "update works set label=?, quantity=?, unit=?, unitprice=?, currency=? where id=?"
	db := store.DB

	_, err := db.Exec(qry, w.Label, w.Quantity, w.Unit, w.UnitPrice, w.Currency, w.ID)
	if err != nil {
		return err
	}

	return nil
}
