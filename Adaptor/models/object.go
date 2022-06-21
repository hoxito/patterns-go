package models

type object struct {
	health int
}

func NewObject(health int) (*object, error) {
	return &object{
		health: health,
	}, nil
}
