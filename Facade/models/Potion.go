package models

type Potion struct {
	ID             string
	Name           string
	Amount         int
	Charges        int
	CurrentCharges int
}

type iPotion interface {
	UsePotion(CharacterID string)
	RechargePotion()
}
