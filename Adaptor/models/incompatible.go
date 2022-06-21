package models

type iProjectile interface {
	impact(*object) error
}

type unknownProjectile struct {
	name   string
	weight float64
	speed  float64
}

func (p *unknownProjectile) impact(obj *object) error {
	return p.impactAdapter(obj)
}

func NewUnknownProjectile(name string, weight, speed float64) (*unknownProjectile, error) {
	return &unknownProjectile{
		name:   name,
		weight: weight,
		speed:  speed,
	}, nil
}
