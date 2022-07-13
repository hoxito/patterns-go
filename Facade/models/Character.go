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
