package testdata

import (
	"github.com/innermond/dots"
)

type CompanyRegister struct {
	Company   dots.Company
	Addresses []dots.Address
	Ibans     []dots.Iban
}

var CompanyRegisterValid = []CompanyRegister{
	{
		dots.Company{Longname: "Acme sa", TIN: "asdasdasertt", RN: "fsfiu/dsd"},
		[]dots.Address{{Address: "address", Location: dots.Point{X: 178.0546, Y: 74.23158}}},
		[]dots.Iban{{Iban: "RO57PORL0000250010000101", Bankname: "bank of some"}},
	},
	{
		dots.Company{Longname: "Acme  ONE sa", TIN: "sasa 34 Europe asdasdasertt", RN: "r158fsfiu/dsd"},
		[]dots.Address{{Address: "address as text", Location: dots.Point{X: 18.0546, Y: -74.23158}}},
		[]dots.Iban{{Iban: "RO20RZBR0000060002651722", Bankname: "bank of xsome"}},
	},
}

var AddressValid = [][]dots.Address{
	[]dots.Address{{Address: "asasai sdsds  76dsds", Location: dots.Point{X: 45, Y: 45}}},
	[]dots.Address{{Address: "opyuiyu\n ssdsdsdd 9898", Location: dots.Point{X: -179, Y: 19.254}}},
}
