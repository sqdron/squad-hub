package api

import (
	"github.com/sqdron/squad/service"
	"fmt"
	"github.com/sqdron/squad/util"
	"github.com/sqdron/squad/policy"
)

type ActivationRequest struct {
	Name string
}

func (a *hubController) Activate(r *service.Specification) (*service.Instruction, error) {
	fmt.Println("Activating...")
	fmt.Println(r)
	//TODO: Hardcode for a while
	res := &service.Instruction{ID:util.GenerateString(10), Group:""}
	switch r.Name {
	case "squad.gateway":
		gatewayInstruction(res)
	}

	return res, nil
}

func gatewayInstruction(i *service.Instruction) {
	i.Policy = []*policy.Policy{{ID:1,Request:"squad.authentification.auth.login"}}
}