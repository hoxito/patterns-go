package models

type Item interface {
	Drop()
	SetElement(element int)
}
type Weapon struct {
	Name   string
	Damage int
}

type Armor struct {
	Name  string
	armor int
}

func NewWeapon(name string, damage int, description string, itemLevel int, rarity int) *Weapon {

	return &Weapon{
		Name:   name,
		Damage: damage,
	}
}
