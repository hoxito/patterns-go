package handlers

import "fmt"

type OwnerValidation struct {
	next clanHandler
}

func (r *OwnerValidation) Run(p *ClanRequest) {
	if p.OwnerValidationDone {
		fmt.Println("Owner Validation already done")
		r.next.Run(p)
		return
	}
	fmt.Println("Owner registering player")
	p.OwnerValidationDone = true
	r.next.Run(p)
}

func (r *OwnerValidation) SetNext(next clanHandler) {
	r.next = next
}
