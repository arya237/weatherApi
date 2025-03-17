package main

import (
	"log"
	"weatherApi/api"
)

func main(){
	
	server := api.NewServer()
	log.Fatal(server.StartServer())
}