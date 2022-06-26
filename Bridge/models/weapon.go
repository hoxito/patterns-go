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
	Imbued  interface{}
}

type Armour struct {
	Name    string
	Defense int
	Element iElement
}

func (w *Weapon) ImbueItem() {
	fmt.Println("Imbueing a weapon")
	w.Element.Imbue()
}
func (w *Weapon) SetElement(e iElement) {
	w.Element = e
}

func (a *Armour) ImbueItem() {
	fmt.Println("Imbueing an Armour")
	a.Element.Imbue()

}
func (a *Armour) SetElement(e iElement) {
	a.Element = e
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
