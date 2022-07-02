package models

type Component interface {
	AddParent(Component)
	GetName() string
	GetParent() Component
	Search(keyword string) []Component
}

type ComponentStruct struct {
	Name   string
	Parent Component
}

func (c *ComponentStruct) AddParent(cmp Component) {
	c.Parent = cmp
}
