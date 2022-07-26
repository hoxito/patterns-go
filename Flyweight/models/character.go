package models

type iCharacter interface {
	setAttackDamage(damage int)
	setAttackDistance(distance int)
	setWeapon(Weapon string)
	getAttackDistance() int
	getWeapon() string
}

type mage struct {
	AttackDamage   int    `json:"attackDamage"`
	AttackDistance int    `json:"attackDistance"`
	Weapon         string `json:"weapon"`
}

func (c *mage) setWeapon(Weapon string) {
	c.Weapon = Weapon
}

func (c *mage) getWeapon() string {
	return c.Weapon
}
func (c *mage) setAttackDamage(damage int) {
	c.AttackDamage = damage
}

func (c *mage) getAttackDamage() int {
	return c.AttackDamage
}

func (c *mage) setAttackDistance(damage int) {
	c.AttackDistance = damage
}

func (c *mage) getAttackDistance() int {
	return c.AttackDistance
}

func newMage() iCharacter {
	return &mage{
		AttackDamage:   80,
		AttackDistance: 200,
		Weapon:         "Staff",
	}
}

type warrior struct {
	AttackDamage   int    `json:"attackDamage"`
	AttackDistance int    `json:"attackDistance"`
	Weapon         string `json:"weapon"`
}

func (c *warrior) setWeapon(Weapon string) {
	c.Weapon = Weapon
}

func (c *warrior) getWeapon() string {
	return c.Weapon
}
func (c *warrior) setAttackDamage(damage int) {
	c.AttackDamage = damage
}

func (c *warrior) getAttackDamage() int {
	return c.AttackDamage
}

func (c *warrior) setAttackDistance(damage int) {
	c.AttackDistance = damage
}

func (c *warrior) getAttackDistance() int {
	return c.AttackDistance
}
func newWarrior() iCharacter {
	return &warrior{
		AttackDamage:   150,
		AttackDistance: 50,
		Weapon:         "sword",
	}
}
