package characters

type orc struct {
}

// each character class varies its attributes depending on the race
func (a *orc) CreateWarrior(name string) iWarrior {
	return &OrcWarrior{
		Warrior: Warrior{
			character: character{
				Name:           "orc warrior " + name,
				AttackDamage:   200,
				AttackDistance: 60,
				Weapon:         "Axe",
			},
			Warcrys: []string{"AAAAAA"},
		},
	}
}

func (a *orc) CreateMage(name string) iMage {
	return &OrcMage{
		Mage: Mage{
			character: character{
				Name:           "Shaman " + name,
				AttackDamage:   80,
				AttackDistance: 200,
				Weapon:         "Scepter",
			},
			Spells: []string{"fireball"},
		},
	}
}
