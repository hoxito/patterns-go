package main

type object struct {
	health int
}

func newObject(health int) (*object, error) {
	return &object{
		health: health,
	}, nil
}
