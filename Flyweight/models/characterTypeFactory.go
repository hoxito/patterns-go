package models

import "fmt"

const (
	mageType    = "mage"
	warriorType = "warrior"
)

var (
	characterTypeFactorySingleton = &CharacterTypeFactory{
		CharacterTypeMap: make(map[string]iCharacterType),
	}
)

type CharacterTypeFactory struct {
	CharacterTypeMap map[string]iCharacterType
}

func (d *CharacterTypeFactory) GetCharacterByType(CharacterType string) (iCharacterType, error) {
	if d.CharacterTypeMap[CharacterType] != nil {
		return d.CharacterTypeMap[CharacterType], nil
	}

	if CharacterType == mageType {
		d.CharacterTypeMap[CharacterType] = newMage()
		return d.CharacterTypeMap[CharacterType], nil
	}
	if CharacterType == warriorType {
		d.CharacterTypeMap[CharacterType] = newWarrior()
		return d.CharacterTypeMap[CharacterType], nil
	}

	return nil, fmt.Errorf("Wrong character type passed")
}

func GetCharacterTypeFactorySingleton() *CharacterTypeFactory {
	return characterTypeFactorySingleton
}
