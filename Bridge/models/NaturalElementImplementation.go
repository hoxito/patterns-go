package models

import (
	"encoding/json"
	"fmt"
)

type NaturalElement struct {
	Element
	Amount int
}

func NewNaturalElement(name string, suffix string, amount int) *NaturalElement {

	return &NaturalElement{
		Element: Element{
			Name:   name,
			Suffix: suffix,
		},
		Amount: amount,
	}
}

func (ne *NaturalElement) Imbue() {
	fmt.Printf("adding %s damage to weapon\n", ne.Name)

	b, err := json.MarshalIndent(ne, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Imbued: ", string(b))

}
