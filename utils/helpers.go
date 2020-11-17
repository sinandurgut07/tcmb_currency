package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	goCache "github.com/patrickmn/go-cache"
)

func GetEnvVars(key string) string {
	c, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("no `%s`", key))
	}
	return c
}

func GetOrSetCurrencies(cache *goCache.Cache) (*tcmbResponse, error) {
	cacheKey := fmt.Sprintf("currencies")
	if currencies, found := cache.Get(cacheKey); found {
		return currencies.(*tcmbResponse), nil
	}
	return processAllCurrencies()
}


func processAllCurrencies() (*tcmbResponse, error) {
	r, err := processHTTPRequest("GET", BaseURL, AllCurrenciesEndpoint)
	if err != nil {
		return nil, err
	}
	resp := new(tcmbResponse)
	if err := xml.Unmarshal(r, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func processHTTPRequest(method, url, endpoint string) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", url, endpoint), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
