package characters

type iCharacter interface {
	setName(name string)
	setAttackDamage(damage int)
	setAttackDistance(distance int)
	setWeapon(weapon string)
	getAttackDistance() int
	getWeapon() string
	getName() string
}

type character struct {
	name           string
	attackDamage   int
	attackDistance int
	weapon         string
}

func (c *character) setName(name string) {
	c.name = name
}

func (c *character) getName() string {
	return c.name
}

func (c *character) setWeapon(weapon string) {
	c.weapon = weapon
}

func (c *character) getWeapon() string {
	return c.weapon
}
func (c *character) setAttackDamage(damage int) {
	c.attackDamage = damage
}

func (c *character) getAttackDamage() int {
	return c.attackDamage
}

func (c *character) setAttackDistance(damage int) {
	c.attackDistance = damage
}

func (c *character) getAttackDistance() int {
	return c.attackDistance
}
