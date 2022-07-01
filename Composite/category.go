package main

import "fmt"

type Category struct {
	ComponentStruct
	Components []Component
	Name       string
	Parent     Component
}

func (c *Category) getName() string {
	return c.Name
}
func (c *Category) getParent() Component {
	return c.Parent
}
func (c *Category) hasParent() bool {
	if c.Name == "" {
		return false
	}
	return true
}
func (c *Category) getParents() (result string) {

	if c.Parent == nil {
		parentName := c.Parent.getName()
		return fmt.Sprintf("%s, %s", result, parentName)

		return ""
	}
}

func (c *Category) add(cmp Component) {
	cmp.addParent(c)
	c.Components = append(c.Components, cmp)

}
