package main

import "fmt"

type item struct {
	Id   string
	Name string
	Parent Component
}

func (i *item) getParent() {
	return i.Parent
}
func (i *item) Drop() {
	fmt.Printf("Dropping %s", i.Name)
}

func (i *item) getName() string {
	return i.name
}
