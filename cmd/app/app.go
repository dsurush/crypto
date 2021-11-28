package app

import (
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
	"net/http"
)
type MainServer struct {
	pool *gorm.DB
	router *httprouter.Router
}

func NewMainServer(pool *gorm.DB, router *httprouter.Router) *MainServer {
	return &MainServer{pool: pool, router: router}
}


func (server *MainServer) Start() {
	server.InitRouts()
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
