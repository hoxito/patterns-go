package models

type Character struct {
	CharacterType iCharacterType
	Type          string
	status        string
	lat           int
	long          int
	speed         float32
}

func newPlayer(CharacterType, Type string) *Character {
	charType, _ := getCharacterTypeFactorySingleton().getCharacterByType(CharacterType)
	return &Character{
		CharacterType: charType,
		Type:          CharacterType,
	}
}

func (c *Character) newLocation(lat, long int) {
	c.lat = lat
	c.long = long
}

func (c *Character) newStatus(status string) {
	c.status = status
}

func (c *Character) newSpeed(speed float32) {
	c.speed = speed
}
