package characters

type character struct {
	name           string
	attackDamage   int
	attackDistance int
	weapon         string
}

func (c *character) getCharacter() character {
	return character{
		name:           c.name,
		attackDamage:   c.attackDamage,
		attackDistance: c.attackDistance,
		weapon:         c.weapon,
	}
}
