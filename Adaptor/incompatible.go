package main

type iProjectile interface {
	impact(*object) error
}

type unknownProjectile struct {
	name   string
	weight float64
	speed  float64
}

func (p *unknownProjectile) impact(*object) error {
	return nil
}

func newUnknownProjectile(name string, weight, speed float64) (*unknownProjectile, error) {
	return &unknownProjectile{
		name:   name,
		weight: weight,
		speed:  speed,
	}, nil
}
