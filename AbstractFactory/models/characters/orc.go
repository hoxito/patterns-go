package characters

type orc struct {
}

// each character class varies its attributes depending on the race
func (a *orc) createWarrior(name string) iWarrior {
	return &OrcWarrior{
		Warrior: Warrior{
			character: character{
				name:           "orc warrior " + name,
				attackDamage:   200,
				attackDistance: 60,
				weapon:         "Axe",
			},
			warcrys: []string{"Fus", "Fus Ro"},
		},
	}
}

func (a *orc) createMage(name string) iMage {
	return &OrcMage{
		Mage: Mage{
			character: character{
				name:           "Shaman " + name,
				attackDamage:   80,
				attackDistance: 200,
				weapon:         "Scepter",
			},
			spells: []string{"fireball"},
		},
	}
}
