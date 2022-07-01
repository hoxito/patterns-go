package main

type Component interface {
	addParent(Component)
	getName() string
	getParent() Component
}

type ComponentStruct struct {
}

func (c *ComponentStruct) addParent(cmp Component) {

}
