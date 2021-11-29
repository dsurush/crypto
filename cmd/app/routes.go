package app

import "log"

func (server *MainServer) InitRouts() {
	log.Printf("init routes....\n")
	server.router.GET("/service/price", server.Test)

}
