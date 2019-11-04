package testdata

import (
	"math/big"

	"github.com/innermond/dots"
)

var WorkValid = []dots.Work{
	{Label: "Acme sa", Quantity: 12.48, Unit: dots.WorkUnit("buc"), UnitPrice: *big.NewRat(234, 100), Currency: dots.Currency("ron")},
}
