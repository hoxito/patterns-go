package models

type weapon struct {
	Name        string
	Damage      int
	Description string
	ItemLevel   int
	Modifiers   [4]weaponModifier
}
