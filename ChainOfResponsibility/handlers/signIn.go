package handlers

import "fmt"

type SignIn struct {
	next ClanHandler
}

func (r *SignIn) Run(c *ClanRequest) {
	if c.SignInDone {
		fmt.Println("SignIn  already done")
		r.next.Run(c)
		return
	}
	fmt.Println("Signing In  player")
	c.SignInDone = true
	r.next.Run(c)
}

func (r *SignIn) SetNext(next ClanHandler) {
	r.next = next
}
