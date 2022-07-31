package main

import (
	"github.com/hoxyto/patterns-go/Flyweight/models"
)

type Character = models.Character

type game struct {
	mages    []*Character
	warriors []*Character
}

func newGame() *game {
	return &game{
		mages:    make([]*Character, 1),
		warriors: make([]*Character, 1),
	}
}

func (c *game) addMage(Type string) {
	player := models.NewCharacter("Mage", Type)
	c.mages = append(c.mages, player)
	return
}

func (c *game) addWarrior(Type string) {
	player := models.NewCharacter("Warrior", Type)
	c.warriors = append(c.warriors, player)
	return
}
