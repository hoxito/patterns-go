package handlers

import (
	"github.com/hoxyto/patterns-go/ChainOfResponsibility/models"
)

type Character = models.Character

type ClanRequest struct {
	Character                  *Character
	OwnerValidationDone        bool
	SeniorMemberValidationDone bool
	SignInDone                 bool
}

type ClanHandler interface {
	Run(*ClanRequest)
	SetNext(ClanHandler)
}
