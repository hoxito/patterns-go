package main

import "fmt"

type client struct {
}

func (c *client) ImpactProjectileWithObject(p iProjectile, o *object) {
	fmt.Println("Client impacts projectile.")
	p.impact(o)
}
func newClient() *client {
	return &client{}
}
