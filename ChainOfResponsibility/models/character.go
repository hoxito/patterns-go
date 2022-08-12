package models

type Character struct {
	Name         string
	AttackDamage int
	Level        int
}

func (c *Character) setName(Name string) {
	c.Name = Name
}

func (c *Character) getName() string {
	return c.Name
}

func (c *Character) setAttackDamage(damage int) {
	c.AttackDamage = damage
}

func (c *Character) getAttackDamage() int {
	return c.AttackDamage
}

func (c *Character) setLevel(damage int) {
	c.Level = damage
}

func (c *Character) getLevel() int {
	return c.Level
}

func NewCharacter(name string) (*Character, error) {

	return &Character{
		Name:         name,
		AttackDamage: 100,
		Level:        1,
	}, nil
}
