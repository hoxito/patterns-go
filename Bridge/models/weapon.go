package models

import "fmt"

type Item interface {
	ImbueItem()
	SetElement(element Element)
}
type Weapon struct {
	Name    string
	Damage  int
	Element iElement
}

type Armour struct {
	Name    string
	Defense int
	Element iElement
}

func (i *Weapon) ImbueItem() {
	fmt.Println("Imbueing a weapon")
	i.Element.Imbue()
}
func (i *Weapon) SetElement(e iElement) {
	i.Element = e
}

func (i *Armour) ImbueItem() {
	fmt.Println("Imbueing an Armour")
	i.Element.Imbue()

}
func (i *Armour) SetElement(e iElement) {
	i.Element = e
}
func NewWeapon(name string, damage int) *Weapon {

	return &Weapon{
		Name:   name,
		Damage: damage,
	}
}
func NewArmour(name string, defense int) *Armour {

	return &Armour{
		Name:    name,
		Defense: defense,
	}
}
