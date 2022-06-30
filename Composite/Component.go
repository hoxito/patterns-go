package main

type Component interface {
	addParent(Component)
	getName() string
	getParent() Component
}
