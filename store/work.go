package store

import (
	"database/sql/driver"
	"math/big"

	"github.com/innermond/dots"
)

type workOp struct {
}

type Rational big.Rat

// valuer implementation
func (r Rational) Value() (driver.Value, error) {
	br := big.Rat(r)
	return br.FloatString(2), nil
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

	res, err := db.Exec(qry, w.Label, w.Quantity, w.Unit, Rational(w.UnitPrice), w.Currency)
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

/*
func (op *workOp) Modify(c dots.Work) error {
	qry := "update companies set longname=?, tin=?, rn=?, is_client=?, is_contractor=? where id=?"
	db := store.DB

	_, err := db.Exec(qry, c.Longname, c.TIN, c.RN, c.IsClient, c.IsContractor, c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (op *workOp) Register(c dots.Work, addrr []dots.Address, ibans []dots.Iban) (int, error) {
	qryWork := "insert into companies (longname, tin, rn, is_client, is_contractor) values(?, ?, ?, ?, ?)"
	qryAddresses := "insert into work_addresses (work_id, address, location) values"
	qryIbans := "insert into work_ibans (work_id, iban, bankname) values"
	db := store.DB

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	resWork, err := tx.Exec(qryWork, c.Longname, c.TIN, c.RN, c.IsClient, c.IsContractor)
	if err != nil {
		return 0, err
	}
	cid, err := resWork.LastInsertId()
	if err != nil {
		return 0, err
	}

	aa := []string{}
	xx := []interface{}{}
	for _, adr := range addrr {
		aa = append(aa, "(?, ?, ST_SRID(Point(?, ?), 4326))")
		xx = append(xx,
			[]interface{}{
				cid,
				adr.Address,
				adr.Location.X,
				adr.Location.Y,
			}...,
		)
	}
	qryAddresses += " " + strings.Join(aa, ",")
	_, err = tx.Exec(qryAddresses, xx...)
	if err != nil {
		return 0, err
	}

	aa = []string{}
	xx = []interface{}{}
	for _, iban := range ibans {
		aa = append(aa, "(?, ?, ?)")
		xx = append(xx,
			[]interface{}{
				cid,
				iban.Iban,
				iban.Bankname,
			}...,
		)
	}
	qryIbans += " " + strings.Join(aa, ",")
	_, err = tx.Exec(qryIbans, xx...)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(cid), err
}
*/
