package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Flyweight/models"
)

func main() {
	game := newGame()

	//Add Mage
	game.addMage("mage")
	game.addMage("mage")
	game.addMage("mage")
	game.addMage("mage")

	//Add Warrior
	game.addWarrior("warrior")
	game.addWarrior("warrior")
	game.addWarrior("warrior")

	characterTypeFactoryInstance := models.GetCharacterTypeFactorySingleton()

	for characterType, Type := range characterTypeFactoryInstance.CharacterTypeMap {
		fmt.Printf("character type : %s\ntype attack damage: %d\ntype attack distance: %d\n type weapon: %s\n", characterType, Type.GetAttackDamage(), Type.GetAttackDistance(), Type.GetWeapon())
	}
}
