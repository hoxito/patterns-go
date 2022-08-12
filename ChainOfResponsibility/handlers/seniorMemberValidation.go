package handlers

import "fmt"

type SeniorMemberValidation struct {
	next ClanHandler
}

func (r *SeniorMemberValidation) Run(c *ClanRequest) {
	if c.SeniorMemberValidationDone {
		fmt.Println("Senior Member Validation already done")
		r.next.Run(c)
		return
	}
	fmt.Println("Senior Member registering player")
	c.SeniorMemberValidationDone = true
	r.next.Run(c)
}

func (r *SeniorMemberValidation) SetNext(next ClanHandler) {
	r.next = next
}
