package utils

import "encoding/xml"


type tcmbResponse struct {
	XMLName  xml.Name `xml:"Tarih_Date" json:"Tarih_Date"`
	Text     string   `xml:"chardata" json:"chardata"`
	Tarih    string   `xml:"Tarih,attr" json:"Tarih,attr"`
	Date     string   `xml:"Date,attr" json:"Date,attr"`
	BultenNo string   `xml:"Bulten_No,attr" json:"Bulten_No,attr"`
	Currency []Currency `xml:"Currency" json:"Currency"`
} 

type Currency struct {
	Text            string `xml:"chardata" json:"chardata"`
	CrossOrder      string `xml:"CrossOrder,attr" json:"CrossOrder,attr"`
	Kod             string `xml:"Kod,attr" json:"Kod,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr" json:"CurrencyCode,attr"`
	Unit            string `xml:"Unit" json:"Unit"`
	Isim            string `xml:"Isim" json:"Isim"`
	CurrencyName    string `xml:"CurrencyName" json:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying" json:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling" json:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying" json:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling" json:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD" json:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther" json:"CrossRateOther"`
} 