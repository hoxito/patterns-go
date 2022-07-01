package main

func main() {
	RustySword := &item{
		ComponentStruct: ComponentStruct{
			Name: "Rusty sword",
		},
	}
	healthPotion := &item{ComponentStruct: ComponentStruct{Name: "minor healing potion"}}

	weapons := &Category{
		ComponentStruct: ComponentStruct{
			Name: "Weapons",
		},
	}

	weapons.add(RustySword)

	Potions := &Category{
		ComponentStruct: ComponentStruct{
			Name: "Potions",
		},
	}
	weapons.add(RustySword)
	Potions.add(healthPotion)
	weapons.add(Potions)

	weapons.search("min")
}
