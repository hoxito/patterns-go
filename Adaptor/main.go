package main

import "fmt"

func main() {
	cli := newClient()
	projectile, _ := newProjectile("arrow", 100, 280)
	unknownProjectile, _ := newUnknownProjectile("floating arrow", 100.90, 280.22)
	var object = &object{
		Health: 100000,
	}
	fmt.Println("object:", object)
	cli.ImpactProjectileWithObject(projectile, object)
	fmt.Println("object:", object)
	cli.ImpactProjectileWithObject(unknownProjectile, object)
	fmt.Println("object:", object)

}
