package handlers

import (
	"github.com/hoxyto/patterns-go/ChainOfResponsibility/models"
)

type Character = models.Character

type ClanHandler interface {
	run(*Character)
	setNext(ClanHandler)
}
