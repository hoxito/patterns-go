package characters

type undead struct {
}

// each character class varies its attributes depending on the race

// each character class varies its attributes depending on the race
func (a *undead) createWarrior(name string) iWarrior {
	return &UndeadWarrior{
		Warrior: Warrior{
			character: character{
				name:           "Elder warrior " + name,
				attackDamage:   200,
				attackDistance: 60,
				weapon:         "Rusty Sword",
			},
			warcrys: []string{"Fus", "Fus Ro"},
		},
	}
}

func (a *undead) createMage(name string) iMage {
	return &UndeadMage{
		Mage: Mage{
			character: character{
				name:           name,
				attackDamage:   80,
				attackDistance: 200,
				weapon:         "Broken Staff",
			},
			spells: []string{"fireball"},
		},
	}
}
