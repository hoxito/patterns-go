package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Flyweight/client"
	"github.com/hoxyto/patterns-go/Flyweight/models"
)

func main() {
	game := client.NewGame()
	fmt.Println("creatimng mages")
	//Add Mage
	game.AddMage("mage")
	game.AddMage("mage")
	game.AddMage("mage")

	fmt.Println("creating warriors")

	//Add Warrior
	game.AddWarrior("warrior")
	game.AddWarrior("warrior")
	game.AddWarrior("warrior")

	characterTypeFactoryInstance := models.GetCharacterTypeFactorySingleton()
	fmt.Println(characterTypeFactoryInstance.CharacterTypeMap)
	for characterType, Type := range characterTypeFactoryInstance.CharacterTypeMap {
		fmt.Printf("character type : %s\ntype attack damage: %d\ntype attack distance: %d\n type weapon: %s\n", characterType, Type.GetAttackDamage(), Type.GetAttackDistance(), Type.GetWeapon())
	}
}
