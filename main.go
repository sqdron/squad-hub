package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad-hub/api"
)

func main() {

	var squad = squad.Client()
	hub:= api.HubController(squad.Options().ApplicationHub)
	squad.Api.Route("activate").Action(hub.Activate)
	squad.RunDetached()
}
