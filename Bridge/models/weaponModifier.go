package models

type iElement interface {
	Imbue()
}

type Element struct {
	Name   string
	Suffix string
}

func NewElement(name string, damage int, description string, suffix string, amount int) *Element {

	return &Element{
		Name:   name,
		Suffix: suffix,
	}
}
