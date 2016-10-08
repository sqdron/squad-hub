package api

import "github.com/sqdron/squad/activation"

type hubController struct{
	endpoint string
}

type IHubController interface{
	Activate(activation.RequestActivation) activation.ServiceInfo
}

func HubController(endpoint string) IHubController{
	return &hubController{endpoint : endpoint}
}

