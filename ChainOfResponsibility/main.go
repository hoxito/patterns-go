package main

import (
	"github.com/hoxyto/patterns-go/ChainOfResponsibility/handlers"
	"github.com/hoxyto/patterns-go/ChainOfResponsibility/models"
)

func main() {

	OwnerApproval := &handlers.OwnerValidation{}

	SeniorApproval := &handlers.SeniorMemberValidation{}
	SeniorApproval.SetNext(OwnerApproval)

	SignIn := &handlers.SignIn{}
	SignIn.SetNext(SeniorApproval)

	var Player, _ = models.NewCharacter("jose")

	PlayerRequest := &handlers.ClanRequest{
		Character:                  *Player,
		OwnerValidationDone:        false,
		SeniorMemberValidationDone: false,
		SignInDone:                 false}

	SignIn.Run(PlayerRequest)
}
