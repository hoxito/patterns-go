package main

import "fmt"

type item struct {
	Id   string
	Name string
}

func (i *item) getName() {
	return i.Name
}
func (i *item) Drop() {
	fmt.Printf("Dropping %s", i.Name)
}

func (i *item) getName() string {
	return f.name
}
