package characters

type iWarrior interface {
	iCharacter
}
type Warrior struct {
	character
	Warcrys []string `json:"Warcrys,omitempty"`
}

func (c *Warrior) learnWarcry(wc string) {
	c.Warcrys = append(c.Warcrys, wc)
}

func (c *Warrior) getWarcrys() []string {
	return c.Warcrys
}
func (c *Warrior) setName(name string) {
	c.character.Name = "dovahkiin " + name
}

// in factory method we would define newWarrior here, but in abstract factory, since
// each class and race combination has its own implementation, we shall do this on each race class.

// return &Warrior{
// 	character: character{
// 		name:           "dovahkiin " + name,
// 		attackDamage:   150,
// 		attackDistance: 50,
// 		weapon:         "sword",
// 	},
// 	Warcrys: []string{"Fus", "Fus Ro"},
// }
