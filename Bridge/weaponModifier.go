package models

type Element struct {
	Name   string
	Suffix string
	Amount int
}

func NewElement(name string, damage int, description string, suffix string, amount int) *Element {

	return &Element{
		Name:   name,
		Suffix: suffix,
		Amount: amount,
	}
}
