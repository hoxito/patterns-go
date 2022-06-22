package models

type Item interface {
	Equip()
	SetModifiers(rarity int)
}
type Weapon struct {
	Name      string
	Damage    int
	Modifiers [4]Modifier
}

type Armor struct {
	Name      string
	armor     int
	Modifiers [4]Modifier
}

func NewWeapon(name string, damage int, description string, itemLevel int, rarity int) *Weapon {
	var Modifiers = make([]Modifier, rarity)

	for i := 0; i < rarity; i++ {
		Modifiers = append(Modifiers, NewModifier())
	}
	return &Weapon{
		Name:      name,
		Damage:    damage,
		Modifiers: Modifiers,
	}
}
