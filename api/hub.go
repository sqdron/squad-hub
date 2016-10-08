package api

import "github.com/sqdron/squad/activation"

type hubController struct{
}

type IHubController interface{
	Activate(activation.RequestActivation) ResponseInstruction
}

func HubController() IHubController{
	return &hubController{}
}

