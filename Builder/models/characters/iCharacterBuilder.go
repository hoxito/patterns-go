package characters

type iCharacterBuilder interface {
	setName(name string)
	setAttackDamage(damage int)
	setAttackDistance(distance int)
	setWeapon(weapon string)
	getAttackDistance() int
	getWeapon() string
	getName() string
	getBuilder() character
}

func getCharacterBuilder(builderType string) iCharacterBuilder {
	if builderType == "warrior" {
		return &warriorCharacterBuilder{}
	}

	if builderType == "mage" {
		return &mageCharacterBuilder{}
	}
	return nil
}
