package main

import (
	"cryptotest/cmd/app"
	"cryptotest/database"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Hello world")
	Start()
}

func Start()  {
	db := database.GetDBConnection()
	router := httprouter.New()
	server := app.NewMainServer(db, router)
	server.Start()
}
