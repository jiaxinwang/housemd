package house

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/levigross/grequests"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// IPInfo ...
type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

// IPAddresses ...
func IPAddresses() (IPInfo, error) {
	var ret IPInfo

	ro := grequests.RequestOptions{
		Params:         map[string]string{},
		RequestTimeout: time.Second * 60,
	}
	if resp, err := grequests.Get(`http://ipinfo.io/json`, &ro); err == nil {
		return ret, json.Unmarshal(resp.Bytes(), &ret)
	} else {
		return ret, err
	}
}
