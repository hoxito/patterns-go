package models

import "fmt"

type Category struct {
	ComponentStruct
	Components []Component
}

func (c *Category) GetName() string {
	return c.Name
}
func (c *Category) GetParent() Component {
	return c.Parent
}
func (c *Category) HasParent() bool {
	if c.Name == "" {
		return false
	}
	return true
}
func (c *Category) Search(keyword string) []Component {
	fmt.Printf("Serching recursively for keyword %s in category %s\n", keyword, c.Name)
	var found = []Component{}
	for _, composite := range c.Components {
		found = append(found, composite.Search(keyword)...)
	}
	return found
}

func (c *Category) GetParents() (result string) {

	if c.Parent == nil {
		parentName := c.Parent.GetName()
		return fmt.Sprintf("%s, %s", result, parentName)

	}
	return ""
}

func (c *Category) Add(cmp Component) {
	cmp.AddParent(c)
	c.Components = append(c.Components, cmp)

}
