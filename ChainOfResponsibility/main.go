package main

import (
	"github.com/hoxyto/patterns-go/ChainOfResponsibility/handlers"
	"github.com/hoxyto/patterns-go/ChainOfResponsibility/models"
)

func main() {

	OwnerApproval := &handlers.OwnerValidation{}

	//Set next for medical department
	SeniorApproval := &handlers.SeniorMemberValidation{}
	SeniorApproval.SetNext(OwnerApproval)

	//Set next for reception department
	SignIn := &handlers.SignIn{}
	SignIn.SetNext(SeniorApproval)

	var Player, _ = models.NewCharacter("jose")

	PlayerRequest := &handlers.ClanRequest{
		Character:                  *Player,
		OwnerValidationDone:        false,
		SeniorMemberValidationDone: false,
		SignInDone:                 false}
	//Patient visiting
	SignIn.Run(PlayerRequest)
}
