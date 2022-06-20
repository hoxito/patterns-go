package main

func (p *unknownProjectile) impactAdapter(obj *object) error {
	adaptedProjectile := &projectile{
		name:   p.name,
		weight: int(p.weight),
		speed:  int(p.speed),
	}
	return adaptedProjectile.impact(obj)

}
