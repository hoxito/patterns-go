package characters

import (
	"fmt"
)

type iCharacterFactory interface {
	CreateWarrior(name string) iWarrior
	CreateMage(name string) iMage
}

func GetCharacterFactory(race string) (iCharacterFactory, error) {
	if race == "human" {
		return &human{}, nil
	}

	if race == "orc" {
		return &orc{}, nil
	}

	if race == "undead" {
		return &undead{}, nil
	}
	return nil, fmt.Errorf("Wrong race type passed")
}
