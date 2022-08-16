package command

import "fmt"

type weapon struct {
	price int32
	Name  string
}

func (w *weapon) Sell() error {
	fmt.Printf("Selling Item: %s", w.Name)
	return nil
}
func (w *weapon) on() {

	fmt.Println("selling Item: %s for price: %d", w.Name, w.price)
}
