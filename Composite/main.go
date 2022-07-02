package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Composite/models"
)

func main() {
	RustySword := &models.Item{
		ComponentStruct: models.ComponentStruct{
			Name: "Rusty sword",
		},
	}
	healthPotion := &models.Item{ComponentStruct: models.ComponentStruct{Name: "minor healing potion"}}

	weapons := &models.Category{
		ComponentStruct: models.ComponentStruct{
			Name: "Weapons",
		},
	}

	weapons.Add(RustySword)

	Potions := &models.Category{
		ComponentStruct: models.ComponentStruct{
			Name: "Potions",
		},
	}
	weapons.Add(RustySword)
	Potions.Add(healthPotion)
	weapons.Add(Potions)

	fmt.Printf("found: %s\n", weapons.Search("min"))
}
