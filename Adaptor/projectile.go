package main

type projectile struct {
	name   string
	weight int
	speed  int
}

func (p *projectile) impact(obj *object) error {

	obj.Health = obj.Health - p.weight*p.speed
	return nil
}

func newProjectile(name string, weight int, speed int) (*projectile, error) {
	return &projectile{
		name:   name,
		weight: weight,
		speed:  speed,
	}, nil
}
