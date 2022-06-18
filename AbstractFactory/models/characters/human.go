package characters

type human struct {
}

// each character class varies its attributes depending on the race

func (a *human) CreateWarrior(name string) iWarrior {
	return &HumanWarrior{
		Warrior: Warrior{
			character: character{
				Name:           "dovahkiin " + name,
				AttackDamage:   150,
				AttackDistance: 50,
				Weapon:         "sword",
			},
			Warcrys: []string{"Fus", "Ro", "Dah"},
		},
	}
}

func (a *human) CreateMage(name string) iMage {
	return &humanMage{
		Mage: Mage{
			character: character{
				Name:           name,
				AttackDamage:   80,
				AttackDistance: 200,
				Weapon:         "Staff",
			},
			Spells: []string{"fireball"},
		},
	}
}
