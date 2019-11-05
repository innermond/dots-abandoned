package dots

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

type Rational big.Rat

func NewRational(a, b int64) Rational {
	return Rational(*big.NewRat(a, b))
}

// valuer implementation
func (r Rational) Value() (driver.Value, error) {
	br := big.Rat(r)
	return br.FloatString(2), nil
}

func (r *Rational) Scan(val interface{}) error {
	if val == nil {
		*r = Rational(big.Rat{})
		return nil
	}

	v, ok := val.([]uint8)
	if !ok {
		return fmt.Errorf("unexpected value %v", val)
	}

	b := &big.Rat{}
	vr, ok := b.SetString(string(v))
	if !ok {
		return fmt.Errorf("unexpected value %s", v)
	}
	*r = Rational(*vr)
	return nil
}

func (r Rational) String() string {
	b := big.Rat(r)
	return b.FloatString(2)
}
