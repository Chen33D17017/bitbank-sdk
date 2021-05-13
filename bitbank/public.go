package bitbank

import (
	"bitbank-sdk/bitbank/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPrice(cryp string) (model.Price, error) {
	var price model.Price
	url := fmt.Sprintf("https://public.bitbank.cc/%s_jpy/ticker", cryp)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return price, fmt.Errorf("getPrice err: %s", err.Error())
	}
	res, err := client.Do(req)
	if err != nil {
		return price, fmt.Errorf("getPrice err: %s", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return price, fmt.Errorf("getPrice err: %s", err.Error())
	}

	err = json.Unmarshal(body, &price)
	if err != nil {
		return price, fmt.Errorf("getPrice err: %s", err.Error())
	}
	return price, nil
}
