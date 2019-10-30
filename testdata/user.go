package testdata

import "github.com/innermond/dots"

type CompanyRegister struct {
	Company   dots.Company
	Addresses []dots.Address
	Ibans     []dots.Iban
}

var CompanyRegisterValid = []CompanyRegister{
	{
		dots.Company{"Acme sa", "asdasdasertt", "fsfiu/dsd"},
		[]dots.Address{{"address", dots.Point{178.0546, 74.23158}}},
		[]dots.Iban{{"234233243iban", "bank of some"}},
	},
	{
		dots.Company{"Acme  ONE sa", "sasa 34 Europe asdasdasertt", "r158fsfiu/dsd"},
		[]dots.Address{{"address as text", dots.Point{18.0546, -74.23158}}},
		[]dots.Iban{{"234232a33243iban", "bank of xsome"}},
	},
}
