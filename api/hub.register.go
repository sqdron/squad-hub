package api

import (
	"fmt"
	"github.com/sqdron/squad/service"
)

func (a *hubController) Register(r service.Specification) bool {
	fmt.Println("Registering...")
	fmt.Println(r)
	return true
}

