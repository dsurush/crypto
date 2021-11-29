package app

import (
	"cryptotest/pkg/service"
	"log"
	"net/http"
)

func (server *MainServer) InitRouts() {
	server.WorkerService.GenerateCombination()
	server.router.GET("/api/test", server.Test)
	log.Println(http.ListenAndServe(":8888", server))
}

func Test()  {
	service.DoRequest("BTC", "USD")
	//service.GenerateCombination()
	//service.Worker(2*time.Second, true, service.SendHello)
}