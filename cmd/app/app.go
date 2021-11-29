package app

import (
	"cryptotest/pkg/service"
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type MainServer struct {
	pool          *sql.DB
	router        *httprouter.Router
	WorkerService *service.Svc
}

func NewMainServer(pool *sql.DB, router *httprouter.Router, workerService *service.Svc) *MainServer {
	return &MainServer{pool: pool, router: router, WorkerService: workerService}
}

func (server *MainServer) Start() {
	server.InitRouts()
	go service.Worker(15*time.Second, true, server.WorkerService.GenerateCombination)
	fmt.Println("a")
	panic(http.ListenAndServe(":8888", server))
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
