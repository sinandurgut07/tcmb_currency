package model

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type Root struct {
	Currencies     []Currency `xml:"Currency" json:"currencies"`
	Date           wrapTime   `xml:"Date,attr" json:"date"`
	BulletinNumber string     `xml:"Bulten_No,attr" json:"bulletin_number"`
}

//NewRoot return new Root object with current currencies
func NewRoot() Root {
	resp, err := http.Get("http://www.tcmb.gov.tr/kurlar/today.xml")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	root := Root{}
	xml.Unmarshal(body, &root)
	return root
}

//GetCurrenciesMap return currencies as map
func (r Root) GetCurrenciesMap() map[string]Currency {
	m := make(map[string]Currency)
	for _, item := range r.Currencies {
		m[item.Code] = item
	}
	return m
}

//Currency structure covers the xml data of the TCMB.
type Currency struct {
	Unit            string  `json:"unit"`
	Name            string  `xml:"Isim" json:"name"`
	CurrencyName    string  `json:"currency_name"`
	ForexBuying     float64 `json:"forex_buying"`
	ForexSelling    float64 `json:"forex_selling"`
	BanknoteBuying  float64 `json:"banknote_buying"`
	BanknoteSelling float64 `json:"banknote_selling"`
	CrossRateUSD    float64 `json:"cross_rate_usd"`
	CrossRateOther  float64 `json:"cross_rate_other"`
	CrossOrder      string  `xml:"CrossOrder,attr" json:"cross_order"`
	Code            string  `xml:"Kod,attr" json:"code"`
	CurrencyCode    string  `xml:"CurrencyCode,attr" json:"currency_code"`
}

const dateFormat string = "01/02/2006"

type wrapTime struct {
	time.Time
}

func (t wrapTime) MarshalText() (result []byte, err error) {
	fmted := t.Format(dateFormat)
	return []byte(fmted), nil
}

func (t *wrapTime) UnmarshalText(text []byte) error {
	parse, err := time.Parse(dateFormat, string(text))
	if err != nil {
		return err
	}
	*t = wrapTime{parse}
	return nil
}
