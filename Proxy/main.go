package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Proxy/models/characters"
)

func main() {

	Char1 := characters.NewProxyCharacter("red", 100)
	Char2 := characters.NewProxyCharacter("blue", 100)

	Char1.Kill(Char2)

	fmt.Printf("------------ Character 1 ------------\n")
	fmt.Printf("Kills: %d\n", Char1.KillCount)
	fmt.Printf("Deaths: %d\n", Char1.DeathCount)
	fmt.Printf("name: %s\n", Char1.Character.Name)
	fmt.Printf("Health: %d\n", Char1.Character.Health)

	fmt.Printf("------------ Character 2 ------------\n")
	fmt.Printf("Kills: %d\n", Char2.KillCount)
	fmt.Printf("Deaths: %d\n", Char2.DeathCount)
	fmt.Printf("name: %s\n", Char2.Character.Name)
	fmt.Printf("Health: %d\n", Char2.Character.Health)

}
