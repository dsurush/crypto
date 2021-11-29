package models

type Currency struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	LASTUPDATE      int64   `json:"LASTUPDATE"`
	SUPPLY          float64     `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}

type Display struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      string `json:"OPEN24HOUR"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO"`
	LOW24HOUR       string `json:"LOW24HOUR"`
	HIGH24HOUR      string `json:"HIGH24HOUR"`
	PRICE           string `json:"PRICE"`
	FROMSYMBOL      string `json:"FROMSYMBOL"`
	TOSYMBOL        string `json:"TOSYMBOL"`
	LASTUPDATE      string   `json:"LASTUPDATE"`
	SUPPLY          string     `json:"SUPPLY"`
	MKTCAP          string `json:"MKTCAP"`
}

type Resp struct {
	RAW     map[string]map[string]Currency `json:"RAW"`
	DISPLAY map[string]map[string]Display  `json:"DISPLAY"`
}
