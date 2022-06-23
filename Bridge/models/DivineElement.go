package models

import (
	"encoding/json"
	"fmt"
)

type DivineElement struct {
	Element
	Amount int
}

func NewDivineElement(name string, suffix string, amount int) *DivineElement {

	return &DivineElement{
		Element: Element{
			Name:   name,
			Suffix: suffix,
		},
		Amount: amount,
	}
}

func (de *DivineElement) Imbue() {
	fmt.Printf("adding incantation: %s to weapon\n", de.Name)

	b, err := json.MarshalIndent(de, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("imbued: ", string(b))
}
