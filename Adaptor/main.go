package main

import (
	"fmt"

	"github.com/hoxyto/patterns-go/Adaptor/models"
)

func main() {

	cli := models.NewClient()
	projectile, _ := models.NewProjectile("arrow", 100, 280)
	unknownProjectile, _ := models.NewUnknownProjectile("floating arrow", 100.90, 280.22)
	object, _ := models.NewObject(10000000)
	fmt.Println("object:", object)
	cli.ImpactProjectileWithObject(projectile, object)
	fmt.Println("object:", object)
	cli.ImpactProjectileWithObject(unknownProjectile, object)
	fmt.Println("object:", object)
}
