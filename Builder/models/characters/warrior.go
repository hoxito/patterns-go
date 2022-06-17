package characters

type warriorBuilder struct {
	character
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

func (c *warriorBuilder) setName(name string) {
	c.character.name = "dovahkiin " + name
}
