package models

import "fmt"

const (
	mageType    = "mage"
	warriorType = "warrior"
)

var (
	characterTypeFactorySingleton = &characterTypeFactory{
		characterTypeMap: make(map[string]CharacterType),
	}
)

type characterTypeFactory struct {
	characterTypeMap map[string]CharacterType
}

func (d *CharacterTypeFactory) getDressByType(CharacterType string) (Character, error) {
	if d.dressMap[CharacterType] != nil {
		return d.dressMap[CharacterType], nil
	}

	if CharacterType == mageType {
		d.dressMap[CharacterType] = newMage()
		return d.dressMap[CharacterType], nil
	}
	if CharacterType == warriorType {
		d.dressMap[CharacterType] = newWarrior()
		return d.dressMap[CharacterType], nil
	}

	return nil, fmt.Errorf("Wrong character type passed")
}

func getCharacterTypeFactorySingleton() *CharacterTypeFactory {
	return characterTypeFactorySingleton
}
