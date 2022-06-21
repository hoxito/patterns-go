package models

import "fmt"

type client struct {
}

func (c *client) ImpactProjectileWithObject(p iProjectile, o *object) {
	fmt.Println("Client impacts projectile.")
	p.impact(o)
}
func NewClient() *client {
	return &client{}
}
