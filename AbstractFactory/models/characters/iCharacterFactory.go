package characters

import "fmt"

type iCharacterFactory interface {
	createWarrior(name string) iWarrior
	createMage(name string) iMage
}

func getSportsFactory(race string) (iCharacterFactory, error) {
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
