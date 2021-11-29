package main

import (
	"cryptotest/cmd/app"
	"cryptotest/database"
	"cryptotest/pkg/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Hello world")
	Start()
}

func Start() {
	db := database.GetDBConnection()
	defer db.Close()
	router := httprouter.New()
	service := service.NewSvc(db)
	server := app.NewMainServer(db, router, service)
	server.Start()
}
