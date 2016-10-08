package api

import "github.com/sqdron/squad/activation"

func (a *hubController) Activate(r activation.RequestActivation) activation.ResponceActivation {
	return activation.ResponceActivation{ID:r.ID, Group:""}
}
