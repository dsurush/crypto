package settings

import (
	"encoding/json"
	"io/ioutil"
)

type AppSetting struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	NameDB   string `json:"name_db"`
}

type Syms struct {
	Fsyms []string `json:"fsyms"`
	Tsyms []string `json:"tsyms"`
}

func ReadSettings(filepath string) AppSetting {
	var AppSetting AppSetting
	doc, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(doc, &AppSetting)
	if err != nil {
		panic(err)
	}
	return AppSetting
}

func ReadCurrency(filepath string) Syms {
	var Syms Syms
	doc, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(doc, &Syms)
	if err != nil {
		panic(err)
	}
	return Syms
}