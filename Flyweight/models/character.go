package models

type Character struct {
	ID   string
	Name string
}

func NewCharacter(ID, Name string) *Character {
	return &Character{
		Name: Name,
		ID:   ID,
	}
}
