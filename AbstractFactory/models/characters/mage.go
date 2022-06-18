package characters

type iMage interface {
	iCharacter
}
type Mage struct {
	character
	spells []string
}

func (c *Mage) learnSpells(spell string) {
	c.spells = append(c.spells, spell)
}

func (c *Mage) getSpells() []string {
	return c.spells
}
