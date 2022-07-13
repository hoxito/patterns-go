package models

import "errors"

type Potion struct {
	ID             string
	Name           string
	Amount         int
	Charges        int
	CurrentCharges int
}

type iPotion interface {
	UsePotion(c Character)
	RechargePotion()
}

func (p Potion) UsePotion(character Character) error {
	if !p.hasCharges() {
		return errors.New("no charges left")
	}
	if character.IsFullHealth() {
		return errors.New("already on full health")
	}
	p.heal(character)
	return nil
}
func (p Potion) hasCharges() bool {
	return p.Charges > 0
}
func (p Potion) heal(c Character) {
	c.Health = c.Health + p.Amount
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
}
