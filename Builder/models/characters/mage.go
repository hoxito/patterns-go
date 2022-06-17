package characters

type mageBuilder struct {
	character
}

func (c *mageBuilder) setWeapon() {
	c.weapon = "sword"
}
func (c *mageBuilder) setAttackDamage() {
	c.attackDamage = 150
}

func (c *mageBuilder) setAttackDistance() {
	c.attackDistance = 50
}

func (c *mageBuilder) setName(name string) {
	c.character.name = name + " the mage"
}
