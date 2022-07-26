package characters

type iCharacter interface {
	setName(name string)
	setAttackDamage(damage int)
	setAttackDistance(distance int)
	setWeapon(weapon string)
	getAttackDistance() int
	getAttackDamage() int
	getWeapon() string
	getName() string
}
