package models

type NaturalElement struct {
	Element
	Amount int
}

func NewNaturalElement(name string, damage int, description string, suffix string, amount int) *NaturalElement {

	return &NaturalElement{
		Element: Element{
			Name:   name,
			Suffix: suffix,
		},
		Amount: amount,
	}
}
