package app

func (server *MainServer) InitRouts() {
	server.router.GET("/api/test", server.Test)

}

func Test() {

}
