package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Singleton/config"
)

func main() {
	a := config.GetConfig()
	a.SetDB("localhost:4242")
	fmt.Println(a.GetDB()) // will print second
	b := config.GetConfig()
	fmt.Println(b.GetDB()) // will also print second
}
