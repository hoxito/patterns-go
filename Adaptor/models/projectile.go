package models

type projectile struct {
	name   string
	weight int
	speed  int
}

func (p *projectile) impact(obj *object) error {

	obj.health = obj.health - p.weight*p.speed
	return nil
}

func NewProjectile(name string, weight int, speed int) (*projectile, error) {
	return &projectile{
		name:   name,
		weight: weight,
		speed:  speed,
	}, nil
}
