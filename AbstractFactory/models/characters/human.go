package characters

type human struct {
}

// each character class varies its attributes depending on the race

func check(i iWarrior) {
	i.setAttackDamage(1)
}
func (a *human) createWarrior(name string) iWarrior {
	return &HumanWarrior{
		Warrior: Warrior{
			character: character{
				name:           "dovahkiin " + name,
				attackDamage:   150,
				attackDistance: 50,
				weapon:         "sword",
			},
			warcrys: []string{"Fus", "Fus Ro"},
		},
	}
}

func (a *human) createMage(name string) iMage {
	return &humanMage{
		Mage: Mage{
			character: character{
				name:           name,
				attackDamage:   80,
				attackDistance: 200,
				weapon:         "Staff",
			},
			spells: []string{"fireball"},
		},
	}
}
