package app

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) Test(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	_, err := writer.Write([]byte("hello api"))
	if err != nil {
		log.Print(err)
	}
}
