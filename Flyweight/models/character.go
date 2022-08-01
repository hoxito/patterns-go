package models

type Character struct {
	CharacterType iCharacterType
	Type          string
	status        string
	lat           int
	long          int
	speed         float32
}

func NewCharacter(Type, TypeName string) *Character {
	CharacterType, _ := GetCharacterTypeFactorySingleton().GetCharacterByType(TypeName)
	return &Character{
		Type:          Type,
		CharacterType: CharacterType,
	}
}

func (c *Character) NewLocation(lat, long int) {
	c.lat = lat
	c.long = long
}

func (c *Character) NewStatus(status string) {
	c.status = status
}

func (c *Character) NewSpeed(speed float32) {
	c.speed = speed
}
