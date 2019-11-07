package testdata

import (
	"math"

	"github.com/innermond/dots"
)

var WorkValid = []dots.Work{
	{Label: "Acme sa", Quantity: 12.48, Unit: dots.WorkUnit("buc"), UnitPrice: dots.NewRational(234, 100), Currency: dots.Currency("ron")},
	{Label: "Booya sa", Quantity: 412.4121548, Unit: dots.WorkUnit("buc"), UnitPrice: dots.NewRational(8234, 100), Currency: dots.Currency("eur")},
}

var WorkStagesValid = []dots.WorkStage{
	{Stage: "stage1", Description: "dsds", Ordered: math.MaxInt32 - 1},
	{Stage: "stage2", Description: "dsds", Ordered: math.MaxInt32 - 2},
	{Stage: "stage3", Description: "dsds", Ordered: math.MaxInt32 - 3},
	{Stage: "stage4", Description: "dsds", Ordered: math.MaxInt32 - 4},
}
