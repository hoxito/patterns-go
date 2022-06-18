package characters

type iMage interface {
	iCharacter
}
type Mage struct {
	character
	Spells []string `json:"spells,omitempty"`
}

func (c *Mage) learnSpells(spell string) {
	c.Spells = append(c.Spells, spell)
}

func (c *Mage) getSpells() []string {
	return c.Spells
}
