package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Bridge/models"
)

func main() {

	Divine := models.NewDivineElement("Evil", "Of Darkness", 150)
	Natural := models.NewNaturalElement("fire", "in fire", 150)

	weaponItem := models.NewWeapon("sword", 10)

	weaponItem.SetElement(Divine)
	weaponItem.ImbueItem()
	fmt.Println()

	weaponItem.SetElement(Natural)
	weaponItem.ImbueItem()
	fmt.Println()

	armourItem := models.NewArmour("chestplate", 10)

	armourItem.SetElement(Divine)
	armourItem.ImbueItem()
	fmt.Println()

	armourItem.SetElement(Natural)
	armourItem.ImbueItem()
	fmt.Println()

}
