package models

import "fmt"

type Item interface {
	Drop()
	SetElement(element Element)
}
type Weapon struct {
	Name    string
	Damage  int
	Element Element
}

type Armour struct {
	Name    string
	Defense int
	Element iElement
}

func (i *Weapon) Drop() {
	fmt.Println("Dropping a weapon")
	i.Element.imbue()
}
func (i *Weapon) SetElement(e Element) {
	fmt.Println("Dropping a weapon")
	i.Element = e
}

func (i *Armour) Drop() {
	fmt.Println("Dropping an Armour")

}
func (i *Armour) SetElement(e Element) {
	fmt.Println("Dropping a weapon")
	i.Element = e
}
func NewWeapon(name string, damage int, description string, itemLevel int, rarity int) *Weapon {

	return &Weapon{
		Name:   name,
		Damage: damage,
	}
}
func NewArmour(name string, defense int, description string, itemLevel int, rarity int) *Armour {

	return &Armour{
		Name:    name,
		Defense: defense,
	}
}
