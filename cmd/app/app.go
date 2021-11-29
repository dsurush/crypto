package app

import (
	"cryptotest/pkg/service"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
type MainServer struct {
	pool *sql.DB
	router *httprouter.Router
	WorkerService *service.Svc
}

func NewMainServer(pool *sql.DB, router *httprouter.Router, workerService *service.Svc) *MainServer {
	return &MainServer{pool: pool, router: router, WorkerService: workerService}
}


func (server *MainServer) Start() {
	server.InitRouts()
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
