package models

import "fmt"

type Weapon struct {
	Price int32
	Name  string
}

func (w *Weapon) Sell() error {
	fmt.Printf("Selling Item: %s", w.Name)
	return nil
}
func (w *Weapon) on() {

	fmt.Println("selling Item: %s for price: %d", w.Name, w.Price)
}
