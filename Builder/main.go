package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/FactoryMethod/models/characters"
)

func main() {
	mage1, err := characters.NewCharacter("mage", "Adria")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(mage1)

	warrior1, err := characters.NewCharacter("warrior", "carl")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(warrior1)

}
