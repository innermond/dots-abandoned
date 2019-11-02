package store

import (
	"fmt"
	"strings"

	"github.com/innermond/dots"
)

func (op *companyOp) AddressAdd(cid int, addresses []dots.Address) ([]int, error) {
	qry := "insert into company_addresses (company_id, address, location) values(?, ?, ST_SRID(Point(?, ?), 4326))"
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

	for _, addr := range addresses {
		res, err := tx.Exec(qry, cid, addr.Address, addr.Location.X, addr.Location.Y)
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

func (op *companyOp) AddressesDelete(cid int, ids []int) error {
	lenids := len(ids)
	in := strings.Repeat(",?", lenids)
	in = in[1:]

	qry := fmt.Sprintf("delete from company_addresses where company_id = ? and id in(%s) limit %d", in, lenids)
	db := store.DB

	err := error(nil)
	_, err = db.Exec(qry, cid)
	if err != nil {
		return err
	}
	return nil
}

func (op *companyOp) AddressModify(cid int, addr dots.Address) error {
	qry := "update company_addresses set address=?, location=ST_SRID(Point(?, ?), 4326) where company_id=? and id=?"
	db := store.DB

	_, err := db.Exec(qry, addr.Address, addr.Location.X, addr.Location.Y, cid, addr.ID)
	if err != nil {
		return err
	}

	return nil
}
