package characters

type warriorBuilder struct {
	character
	warcrys []string
}

func (c *warriorBuilder) setWeapon() {
	c.weapon = "sword"
}
func (c *warriorBuilder) setAttackDamage() {
	c.attackDamage = 150
}

func (c *warriorBuilder) setAttackDistance() {
	c.attackDistance = 50
}

func (c *warriorBuilder) learnWarcry(wc string) {
	c.warcrys = append(c.warcrys, wc)
}

func (c *warriorBuilder) getWarcrys() []string {
	return c.warcrys
}

func (c *warriorBuilder) setName(name string) {
	c.character.name = "dovahkiin " + name
}

func newWarrior(name string) iCharacter {
	return &warriorBuilder{
		character: character{
			name:           "dovahkiin " + name,
			attackDamage:   150,
			attackDistance: 50,
			weapon:         "sword",
		},
		warcrys: []string{"Fus", "Fus Ro"},
	}
}
