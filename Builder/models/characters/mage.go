package characters

//each builder has its own set of methods that build a certain type of object with different properties

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
