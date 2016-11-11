package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad-hub/api"
)

type Options struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func main() {
	opts := &Options{}
	hub := squad.Hub(opts)
	h := api.HubController(hub.Options.Hub)
	hub.Api.Register(h.Register)
	hub.Api.Activate(h.Activate)
	hub.Run()
}

//
//func getDB(o *Options) *gorm.DB {
//	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", o.DbHost, o.DbPort, o.DbUser, o.DbName, o.DbPassword)
//	db, err := gorm.Open("postgres", connectionString)
//	if (err != nil) {
//		fmt.Errorf("DBError", err)
//	}
//	return db
//}
