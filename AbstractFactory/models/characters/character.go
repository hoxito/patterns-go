package characters

type iCharacter interface {
	setName(Name string)
	setAttackDamage(damage int)
	setAttackDistance(distance int)
	setWeapon(Weapon string)
	getAttackDistance() int
	getWeapon() string
	getName() string
}

type character struct {
	Name           string `json:"name"`
	AttackDamage   int    `json:"attackDamage"`
	AttackDistance int    `json:"attackDistance"`
	Weapon         string `json:"weapon"`
}

func (c *character) setName(Name string) {
	c.Name = Name
}

func (c *character) getName() string {
	return c.Name
}

func (c *character) setWeapon(Weapon string) {
	c.Weapon = Weapon
}

func (c *character) getWeapon() string {
	return c.Weapon
}
func (c *character) setAttackDamage(damage int) {
	c.AttackDamage = damage
}

func (c *character) getAttackDamage() int {
	return c.AttackDamage
}

func (c *character) setAttackDistance(damage int) {
	c.AttackDistance = damage
}

func (c *character) getAttackDistance() int {
	return c.AttackDistance
}
