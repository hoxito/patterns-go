package characters

type undead struct {
}

// each character class varies its attributes depending on the race

// each character class varies its attributes depending on the race
func (a *undead) CreateWarrior(name string) iWarrior {
	return &UndeadWarrior{
		Warrior: Warrior{
			character: character{
				Name:           "Elder warrior " + name,
				AttackDamage:   200,
				AttackDistance: 60,
				Weapon:         "Rusty Sword",
			},
			Warcrys: []string{"Fus", "Ro"},
		},
	}
}

func (a *undead) CreateMage(name string) iMage {
	return &UndeadMage{
		Mage: Mage{
			character: character{
				Name:           name,
				AttackDamage:   80,
				AttackDistance: 200,
				Weapon:         "Broken Staff",
			},
			Spells: []string{"fireball"},
		},
	}
}
