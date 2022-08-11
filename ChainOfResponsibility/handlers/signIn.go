package handlers

import "fmt"

type SignIn struct {
	next clanHandler
}

func (r *SignIn) run(p *ClanRequest) {
	if p.SignInDone {
		fmt.Println("SignIn  already done")
		r.next.run(p)
		return
	}
	fmt.Println("Signing In  player")
	p.SignInDone = true
	r.next.run(p)
}

func (r *SignIn) setNext(next clanHandler) {
	r.next = next
}
