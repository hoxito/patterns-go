package characters

type mageBuilder struct {
	character
}

func (c *mageBuilder) setWeapon() {
	c.weapon = "staff"
}
func (c *mageBuilder) setAttackDamage() {
	c.attackDamage = 80
}

func (c *mageBuilder) setAttackDistance() {
	c.attackDistance = 200
}

func (c *mageBuilder) setName(name string) {
	c.character.name = name + " the mage"
}
