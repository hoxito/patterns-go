package models

import (
	"fmt"
	"strings"
)

type Item struct {
	ComponentStruct
	Id string
}

func (i *Item) GetParent() Component {
	return i.Parent
}
func (i *Item) Drop() {
	fmt.Printf("Dropping %s", i.Name)
}

func (i *Item) Search(keyword string) []Component {
	fmt.Printf("Searching for keyword %s in Item %s\n", keyword, i.Name)
	if strings.Contains(i.Name, keyword) {
		return []Component{i}
	}
	return nil
}
func (i *Item) GetName() string {
	return i.Name
}
