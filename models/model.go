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
	SUPPLY          int     `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}

type Display struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	FROMSYMBOL      string `json:"FROMSYMBOL"`
	TOSYMBOL        string `json:"TOSYMBOL"`
	LASTUPDATE      int64   `json:"LASTUPDATE"`
	SUPPLY          int     `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}

type Resp struct {
	RAW     map[string]map[string]Currency `json:"RAW"`
	DISPLAY map[string]map[string]Display  `json:"DISPLAY"`
}
