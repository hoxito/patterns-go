package characters

type warrior struct {
	character
	warcrys []string
}

func (c *warrior) learnWarcry(wc string) {
	c.warcrys = append(c.warcrys, wc)
}

func (c *warrior) getWarcrys() []string {
	return c.warcrys
}
func (c *warrior) setName(name string) {
	c.character.name = "dovahkiin " + name
}

func newWarrior(name string) iCharacter {
	return &warrior{
		character: character{
			name:           "dovahkiin " + name,
			attackDamage:   150,
			attackDistance: 50,
			weapon:         "sword",
		},
		warcrys: []string{"Fus", "Fus Ro"},
	}
}
