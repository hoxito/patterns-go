package models

type iElement interface {
	imbue() string
}

type Element struct {
	Name   string
	Suffix string
}

type DivineElement struct {
	Element
	Amount int
}

func NewElement(name string, damage int, description string, suffix string, amount int) *Element {

	return &Element{
		Name:   name,
		Suffix: suffix,
	}
}
