package service

import (
	"cryptotest/models"
	"cryptotest/settings"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

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

func GenerateCombination() {
	currency := settings.ReadCurrency(`..\cryptotest\settings\currency.json`)
	for _, Fsym := range currency.Fsyms{
		for _, Tsym := range currency.Tsyms{
			fmt.Println(Fsym+Tsym)
		}
	}
}

func DoRequest(fsyms, tsyms string) bool {
	client := &http.Client{}
	URL := fmt.Sprintf(`https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s`, fsyms, tsyms)
	Method := `GET`
	req, err := http.NewRequest(Method, URL, nil)
	if err != nil {
		log.Printf("Can't build request\n")
		return false
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request\n")
		return false
	}
	fmt.Println(`Status Code = `, resp.StatusCode)
	if resp.StatusCode == http.StatusOK {

		data := models.Resp{}
		json.NewDecoder(resp.Body).Decode(&data)
		//_ = json.Unmarshal([]byte(str), &data)
		fmt.Println(data)
		return true
	}
	return false
}

func SendHello()  {
	fmt.Println("hello")
}