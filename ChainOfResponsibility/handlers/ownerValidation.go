package handlers

import "fmt"

type OwnerValidation struct {
	next clanHandler
}

func (r *OwnerValidation) run(p *ClanRequest) {
	if p.OwnerValidationDone {
		fmt.Println("Owner Validation already done")
		r.next.run(p)
		return
	}
	fmt.Println("Owner registering player")
	p.OwnerValidationDone = true
	r.next.run(p)
}

func (r *OwnerValidation) setNext(next clanHandler) {
	r.next = next
}
