package app

import (
	"cryptotest/models"
	"cryptotest/pkg/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) Test(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set(contentType, value)
	fsyms := request.URL.Query().Get(`fsyms`)
	tsyms := request.URL.Query().Get(`tsyms`)
	ok, resp := service.DoRequest(fsyms, tsyms)
	if ok {
		err := json.NewEncoder(writer).Encode(resp)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		str := server.WorkerService.GetFromDBCurr(fsyms + tsyms)
		data := models.Resp{}
		err := json.Unmarshal([]byte(str), &data)
		if err != nil {
			log.Printf("Can't Unmarshal in Api err %e\n", err)
			return
		}
		err = json.NewEncoder(writer).Encode(data)
		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}
