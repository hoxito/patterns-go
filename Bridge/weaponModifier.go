package models

type Modifier struct {
	Name   string
	Suffix string
	Amount int
}

func NewModifier(name string, damage int, description string, suffix string, amount int) *Modifier {

	return &Modifier{
		Name:   name,
		Suffix: suffix,
		Amount: amount,
	}
}
