package testdata

import (
	"github.com/innermond/dots"
)

var WorkValid = []dots.Work{
	{Label: "Acme sa", Quantity: 12.48, Unit: dots.WorkUnit("buc"), UnitPrice: dots.NewRational(234, 100), Currency: dots.Currency("ron")},
	{Label: "Booya sa", Quantity: 412.4121548, Unit: dots.WorkUnit("buc"), UnitPrice: dots.NewRational(8234, 100), Currency: dots.Currency("eur")},
}
