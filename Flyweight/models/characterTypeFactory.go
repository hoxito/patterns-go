package models

import "fmt"

const (
	mageType    = "mage"
	warriorType = "warrior"
)

var (
	characterTypeFactorySingleton = &CharacterTypeFactory{
		characterTypeMap: make(map[string]iCharacter),
	}
)

type CharacterTypeFactory struct {
	characterTypeMap map[string]iCharacter
}

func (d *CharacterTypeFactory) getDressByType(CharacterType string) (iCharacter, error) {
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
