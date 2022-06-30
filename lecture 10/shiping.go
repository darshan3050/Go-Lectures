package main

type ShipingCompany interface {
	calculate()
}
type ups struct {
	Amount int
}
type fedex struct {
	Amount int
}
type ekart struct {
	Amount int
}
type Amazon struct {
	Company ShipingCompany
}

func (a *Amazon) calculate(company ShipingCompany) {
	a.Company.calculate()
}
