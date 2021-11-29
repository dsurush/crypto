package app

import (
	"cryptotest/pkg/service"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"log"
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
	log.Printf("Server is starting...\n")
	server.WorkerService.GenerateCombination()
	server.InitRouts()
	go service.Worker(3*time.Minute, true, server.WorkerService.GenerateCombination)
	panic(http.ListenAndServe(":8888", server))
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
