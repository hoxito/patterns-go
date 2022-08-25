package models

import "fmt"

type Button struct {
	Command Command
}
type KeyBind struct {
	Command Command
}

func (b *Button) Press() {
	fmt.Println("pressing a button")
	b.Command.Execute()
}

func (b *KeyBind) Press() {
	fmt.Println("Using a KeyBind")
	b.Command.Execute()
}
