package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Builder/models/characters"
)

func main() {
	warriorBuilder := characters.GetCharacter("warrior")
	mageBuilder := characters.GetCharacter("mage")

	director := characters.NewDirector(warriorBuilder)
	warrior := director.BuildCharacter("John Doe")

	fmt.Printf("Warrior: %v\n", warrior)
	director.SetBuilder(mageBuilder)
	mage := director.BuildCharacter("Carl")

	fmt.Printf("Mage: %v\n", mage)
}
