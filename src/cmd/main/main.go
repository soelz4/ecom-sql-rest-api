package main

import (
	"fmt"
	"log"

	"ecom/src/cmd/api"
	"ecom/src/config"
	"ecom/src/db"
)

func main() {
	database, err := db.NewMySQLStorage(db.GetConfig())
	if err != nil {
		log.Fatal(err)
	}
	db.InitMySQLStorage(database)
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), database)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
