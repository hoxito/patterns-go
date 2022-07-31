package models

type Character struct {
	CharacterType iCharacterType
	Type          string
	status        string
	lat           int
	long          int
	speed         float32
}

func NewCharacter(CharacterType, Type string) *Character {
	charType, _ := GetCharacterTypeFactorySingleton().GetCharacterByType(CharacterType)
	return &Character{
		CharacterType: charType,
		Type:          CharacterType,
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
