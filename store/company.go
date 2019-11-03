package store

import (
	"strings"

	"github.com/innermond/dots"
)

type companyOp struct {
}

var companyOperations *companyOp

func CompanyOp() *companyOp {
	if companyOperations == nil {
		companyOperations = &companyOp{}
	}
	return companyOperations
}

func (op *companyOp) Add(c dots.Company) (int, error) {
	qry := "insert into companies (longname, tin, rn) values(?, ?, ?)"
	db := store.DB

	res, err := db.Exec(qry, c.Longname, c.TIN, c.RN)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (op *companyOp) Modify(c dots.Company) error {
	qry := "update companies set longname=?, tin=?, rn=? where id=?"
	db := store.DB

	_, err := db.Exec(qry, c.Longname, c.TIN, c.RN, c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (op *companyOp) Delete(cid int) error {
	qry := "delete from companies where id = ? limit 1"
	db := store.DB
	err := error(nil)
	_, err = db.Exec(qry, cid)
	if err != nil {
		return err
	}
	return nil
}

func (op *companyOp) Register(c dots.Company, addrr []dots.Address, ibans []dots.Iban) (int, error) {
	qryCompany := "insert into companies (longname, tin, rn) values(?, ?, ?)"
	qryAddresses := "insert into company_addresses (company_id, address, location) values"
	qryIbans := "insert into company_ibans (company_id, iban, bankname) values"
	db := store.DB

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	resCompany, err := tx.Exec(qryCompany, c.Longname, c.TIN, c.RN)
	if err != nil {
		return 0, err
	}
	cid, err := resCompany.LastInsertId()
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
