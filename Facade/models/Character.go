package models

type Character struct {
	ID        string
	Name      string
	Health    int
	MaxHealth int
}

func (c Character) IsFullHealth() bool {
	return c.Health == c.MaxHealth
}
func (c Character) UsePotionOnSelf(p iPotion) {
	p.UsePotion(c)
}

func (c Character) UsePotionOn(p iPotion, char Character) {
	p.UsePotion(char)
}
