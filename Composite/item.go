package main

import (
	"fmt"
	"strings"
)

type item struct {
	ComponentStruct
	Id string
}

func (i *item) getParent() Component {
	return i.Parent
}
func (i *item) Drop() {
	fmt.Printf("Dropping %s", i.Name)
}
func (c *Category) search(keyword string) []Component {
	fmt.Printf("Serching recursively for keyword %s in category %s\n", keyword, c.Name)
	var found = []Component{}
	for _, composite := range c.Components {
		found = append(found, composite.search(keyword)...)
	}
	return found
}

func (i *item) search(keyword string) []Component {
	fmt.Printf("Searching for keyword %s in item %s\n", keyword, i.Name)
	if strings.Contains(i.Name, keyword) {
		return []Component{i}
	}
	return nil
}
func (i *item) getName() string {
	return i.Name
}
