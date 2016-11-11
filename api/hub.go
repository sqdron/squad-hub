package api

type hubController struct{
	endpoint string
}


func HubController(endpoint string) *hubController{
	return &hubController{endpoint : endpoint}
}

