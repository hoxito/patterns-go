package handlers

import "fmt"

type SignIn struct {
	next clanHandler
}

func (r *SignIn) Run(p *ClanRequest) {
	if p.SignInDone {
		fmt.Println("SignIn  already done")
		r.next.Run(p)
		return
	}
	fmt.Println("Signing In  player")
	p.SignInDone = true
	r.next.Run(p)
}

func (r *SignIn) SetNext(next clanHandler) {
	r.next = next
}
