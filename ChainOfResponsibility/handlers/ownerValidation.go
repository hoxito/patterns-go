package handlers

import "fmt"

type OwnerValidation struct {
	next ClanHandler
}

func (r *OwnerValidation) Run(c *ClanRequest) {
	if c.OwnerValidationDone {
		fmt.Println("Owner Validation already done")
		r.next.Run(c)
		return
	}
	fmt.Println("Owner registering player")
	c.OwnerValidationDone = true

}

func (r *OwnerValidation) SetNext(next ClanHandler) {
	r.next = next
}
