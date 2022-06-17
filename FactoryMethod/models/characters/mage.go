package characters

type mage struct {
	character
	spells []string
}

func (c *mage) learnSpells(spell string) {
	c.spells = append(c.spells, spell)
}

func (c *mage) getSpells() []string {
	return c.spells
}

func newMage(name string) iCharacter {
	return &mage{
		character: character{
			name:           name,
			attackDamage:   80,
			attackDistance: 200,
			weapon:         "Staff",
		},
		spells: []string{"fireball"},
	}
}
