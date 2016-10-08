package api

import "github.com/sqdron/squad/activation"

func (a *hubController) Activate(r activation.RequestActivation) activation.ServiceInfo {
	return activation.ServiceInfo{Endpoint:a.endpoint}
}
