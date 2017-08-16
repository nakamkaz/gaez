package gaez

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

const zaifPublicEP = "https://api.zaif.jp/api/1"

type LastPrice struct {
	LastPrice float64 `json:"last_price"`
}

func GetLastPrice(currencyPair string, w http.ResponseWriter, req *http.Request) (last_price LastPrice, err error) {
	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	resp, err := client.Get(zaifPublicEP + "/last_price/" + currencyPair)
	if err != nil {
		return LastPrice{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return LastPrice{}, err
	}
	json.Unmarshal(body, &last_price)
	return last_price, nil
}
