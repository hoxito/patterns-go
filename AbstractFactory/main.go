package main

import (
	"encoding/json"
	"fmt"

	"github.com/hoxyto/patterns-go/AbstractFactory/models/characters"
)

func main() {
	humanFactory, err := characters.GetCharacterFactory("human")
	if err != nil {
		panic(err)
	}
	orcFactory, _ := characters.GetCharacterFactory("orc")
	undeadFactory, _ := characters.GetCharacterFactory("undead")

	humanWarrior := humanFactory.CreateWarrior("Carl")
	humanMage := humanFactory.CreateMage("Harry")

	b, err := json.MarshalIndent(humanWarrior, "", "  ")
	fmt.Println("human warrior: ", string(b))

	b, err = json.MarshalIndent(humanMage, "", "  ")
	fmt.Println("human mage: ", string(b))

	orcWarrior := orcFactory.CreateWarrior("Harl")
	orcMage := orcFactory.CreateMage("voldemort")

	b, err = json.MarshalIndent(orcWarrior, "", "  ")
	fmt.Println("orc warrior: ", string(b))

	b, err = json.MarshalIndent(orcMage, "", "  ")
	fmt.Println("orc mage: ", string(b))

	undeadWarrior := undeadFactory.CreateWarrior("Parl")
	undeadMage := undeadFactory.CreateMage("Dombuldore")

	b, err = json.MarshalIndent(undeadWarrior, "", "  ")
	fmt.Println("undead warrior: ", string(b))

	b, err = json.MarshalIndent(undeadMage, "", "  ")
	fmt.Println("undead mage: ", string(b))

}
