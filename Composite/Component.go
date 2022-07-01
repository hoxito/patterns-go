package main

type Component interface {
	addParent(Component)
	getName() string
	getParent() Component
	search(string) []Component
}

type ComponentStruct struct {
	Name   string
	Parent Component
}

func (c *ComponentStruct) addParent(cmp Component) {
	c.Parent = cmp
}
