package characters

type characterDirector struct {
	builder iCharacterBuilder
}

func NewDirector(b iCharacterBuilder) *characterDirector {
	return &characterDirector{
		builder: b,
	}
}

func (c *characterDirector) SetBuilder(b iCharacterBuilder) {
	c.builder = b
}

//This function sets the character building order, calling each builder's method.
func (c *characterDirector) BuildCharacter(name string) character {
	c.builder.setName(name)
	c.builder.setAttackDamage()
	c.builder.setAttackDistance()
	c.builder.setWeapon()

	return c.builder.getCharacter()
}
