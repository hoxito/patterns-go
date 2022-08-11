package handlers

import "fmt"

type SeniorMemberValidation struct {
	next clanHandler
}

func (r *SeniorMemberValidation) Run(p *ClanRequest) {
	if p.SeniorMemberValidationDone {
		fmt.Println("Senior Member Validation already done")
		r.next.run(p)
		return
	}
	fmt.Println("Senior Member registering player")
	p.SeniorMemberValidationDone = true
	r.next.run(p)
}

func (r *SeniorMemberValidation) SetNext(next clanHandler) {
	r.next = next
}
