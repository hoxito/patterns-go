package models

import "fmt"

const (
	mageType    = "mage"
	warriorType = "warrior"
)

var (
	characterTypeFactorySingleton = &CharacterTypeFactory{
		characterTypeMap: make(map[string]iCharacterType),
	}
)

type CharacterTypeFactory struct {
	characterTypeMap map[string]iCharacterType
}

func (d *CharacterTypeFactory) getCharacterByType(CharacterType string) (iCharacterType, error) {
	if d.characterTypeMap[CharacterType] != nil {
		return d.characterTypeMap[CharacterType], nil
	}

	if CharacterType == mageType {
		d.characterTypeMap[CharacterType] = newMage()
		return d.characterTypeMap[CharacterType], nil
	}
	if CharacterType == warriorType {
		d.characterTypeMap[CharacterType] = newWarrior()
		return d.characterTypeMap[CharacterType], nil
	}

	return nil, fmt.Errorf("Wrong character type passed")
}

func getCharacterTypeFactorySingleton() *CharacterTypeFactory {
	return characterTypeFactorySingleton
}
