package main

type category struct {
	Components []component
	Name       string
	Parent     component
}

func (f *folder) getParent() component {
	return Parent
}
func (c *category) hasParent() bool {
	if c.name == nil {
		return false
	}
	return true
}
func (c *category) getParents() (result string) {

	if c.hasParent {
		return (result + c.Parent)
	}
	return Parent
}

func (c *Category) add(cmp component) {
	cmp.Parent = c
	c.Components = append(c.components, cmp)

}
