package characters

type iCharacterBuilder interface {
	setName(name string)
	setAttackDamage()
	setAttackDistance()
	setWeapon()
	getCharacter() character
}

func GetCharacter(builderType string) iCharacterBuilder {
	if builderType == "warrior" {
		return &warriorBuilder{}
	}

	if builderType == "mage" {
		return &mageBuilder{}
	}
	return nil
}
