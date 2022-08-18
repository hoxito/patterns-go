package main

import "github.com/hoxyto/patterns-go/Command/models"

func main() {
	sword := &models.Weapon{Price: 300,
		Name: "Excalibur"}

	SellCommand := &models.SellCommand{
		Item: sword,
	}

	SellButton := &models.Button{
		Command: SellCommand,
	}
	SellButton.Press()

}
