package service

import (
	"cryptotest/database"
	"cryptotest/models"
	"cryptotest/settings"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Svc struct {
	pool *sql.DB
}

func NewSvc(pool *sql.DB) *Svc {
	return &Svc{pool: pool}
}

type ErrorRes struct {
	Response string `json:"Response"`
}

func (svc *Svc) InsertNewCurr(Curr2Curr, RawDisp string) {
	_, err := svc.pool.Exec(database.InsertInfo, Curr2Curr, RawDisp)
	if err != nil {
		log.Printf("Can't InsertNewCurr %e", err)
	}
}

func (svc *Svc) GetFromDBCurr(Curr2Curr string) string {
	var data string
	_ = svc.pool.QueryRow(database.GetFromDBInfo, Curr2Curr).Scan(&data)
	return data
}

func (svc *Svc) UpdateCurr(Curr2Curr, RawDisp string) {
	_, err := svc.pool.Exec(database.UpdateInfo, RawDisp, Curr2Curr)
	if err != nil {
		log.Printf("Can't Update Curr %e", err)
	}
}

func (svc *Svc) CheckHasCurr(Curr2Curr string) (int, error) {
	var count int
	err := svc.pool.QueryRow(database.CheckHasInfo, Curr2Curr).Scan(&count)
	if err != nil {
		log.Printf("CheckHasCurr Can't scan %e", err)
		return count, err
	}
	return count, nil
}

func (svc *Svc) DoUpdate(fsym, tsym, rawDisplay string) {
	ok, err := svc.CheckHasCurr(fsym + tsym)
	if err != nil {
		return
	}
	if ok == 1 {
		svc.UpdateCurr(fsym+tsym, rawDisplay)
	} else {
		svc.InsertNewCurr(fsym+tsym, rawDisplay)
	}
}

// Worker - воркер который запускает каждые d(duration) функцию f если он еще не запущен
func Worker(d time.Duration, async bool, f func()) {
	var reentranceFlag int64
	for range time.Tick(d) {
		go func() {
			if !async {

				if atomic.CompareAndSwapInt64(&reentranceFlag, 0, 1) {
					defer atomic.StoreInt64(&reentranceFlag, 0)
				} else {
					log.Println("Previous worker in process now")
					return
				}
			}
			f()
		}()
	}
}

func (receiver *Svc) GenerateCombination() {
	currency := settings.ReadCurrency(`..\cryptotest\settings\currency.json`)
	for _, Fsym := range currency.Fsyms {
		for _, Tsym := range currency.Tsyms {
			log.Println(Fsym + Tsym)
			func() {
				ok, resp := DoRequest(Fsym, Tsym)
				if !ok {
					log.Println("Request is bad")
					return
				}
				marshal, err := json.Marshal(resp)
				if err != nil {
					log.Println("can't marshal response")
					return
				}
				receiver.DoUpdate(Fsym, Tsym, string(marshal))
			}()
		}
	}
}


func DoRequest(fsyms, tsyms string) (bool, models.Resp) {
	client := &http.Client{}
	data := models.Resp{}
	URL := fmt.Sprintf(`https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s`, fsyms, tsyms)
	Method := `GET`
	req, err := http.NewRequest(Method, URL, nil)
	if err != nil {
		log.Printf("Can't build request\n")
		return false, data
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request\n")
		return false, data
	}
	if resp.StatusCode == http.StatusOK {
		allData, _ := ioutil.ReadAll(resp.Body)
		var ErrorRes ErrorRes
		err := json.Unmarshal(allData, &ErrorRes)
		if err != nil {
			log.Println("json invalid err = ", err)
			return false, data
		}
		if ErrorRes.Response != ""{
			log.Println("Response is Error", err)
			return false, data
		}
		err = json.Unmarshal(allData, &data)
		if err != nil {
			log.Println("json invalid err = ", err)
			return false, data
		}
		log.Println(data)
		return true, data
	}
	return false, data
}

func SendHello() {
	fmt.Println("hello")
}
