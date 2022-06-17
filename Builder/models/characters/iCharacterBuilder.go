package characters

type iCharacterBuilder interface {
	setName(name string)
	setAttackDamage()
	setAttackDistance()
	setWeapon()
	getCharacter() character
}

func getCharacter(builderType string) iCharacterBuilder {
	if builderType == "warrior" {
		return &warriorBuilder{}
	}

	if builderType == "mage" {
		return &mageBuilder{}
	}
	return nil
}
