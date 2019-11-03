package store

import (
	"fmt"
	"strings"

	"github.com/innermond/dots"
)

func (op *companyOp) IbansAdd(cid int, ibans []dots.Iban) ([]int, error) {
	qry := "insert into company_ibans (company_id, iban, bankname) values(?, ?, ?)"
	db := store.DB

	var (
		inserted []int
		empty    = []int{0}
	)

	tx, err := db.Begin()
	if err != nil {
		return empty, err
	}
	defer tx.Rollback()

	for _, iban := range ibans {
		res, err := tx.Exec(qry, cid, iban.Iban, iban.Bankname)
		if err != nil {
			return empty, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return empty, err
		}
		inserted = append(inserted, int(id))
	}

	err = tx.Commit()
	if err != nil {
		return empty, err
	}

	return inserted, nil
}

func (op *companyOp) IbansDelete(cid int, ids []int) error {
	lenids := len(ids)
	in := strings.Repeat(",?", lenids)
	in = in[1:]

	qry := fmt.Sprintf("delete from company_ibans where company_id = ? and id in(%s) limit %d", in, lenids)
	db := store.DB

	err := error(nil)
	_, err = db.Exec(qry, cid)
	if err != nil {
		return err
	}
	return nil
}

func (op *companyOp) IbanModify(cid int, iban dots.Iban) error {
	qry := "update company_ibans set iban=?, bankname=? where company_id=? and id=? limit 1"
	db := store.DB

	_, err := db.Exec(qry, iban.Iban, iban.Bankname, cid, iban.ID)
	if err != nil {
		return err
	}

	return nil
}
