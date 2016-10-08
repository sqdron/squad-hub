package api

import "github.com/sqdron/squad/activation"

//type RequestActivate struct {
//	ID string
//}

type ResponseInstruction struct {
	Group string
}

func (a *hubController) Activate(r activation.RequestActivation) ResponseInstruction {
	return ResponseInstruction{Group:""}
}
