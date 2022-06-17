package characters

// Character struct is the base struct warrior and mage builders will composite with.
type character struct {
	name           string
	attackDamage   int
	attackDistance int
	weapon         string
}

//getCharacter returns the object after the entire object has been built
func (c *character) getCharacter() character {
	return character{
		name:           c.name,
		attackDamage:   c.attackDamage,
		attackDistance: c.attackDistance,
		weapon:         c.weapon,
	}
}
