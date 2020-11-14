package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sinandurgut07/tcmb_currency/model"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/currency/{currency}", CurrencyHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	currencies := model.NewRoot()
	data := model.APIResponder{
		Code: 200,
		Data: currencies.GetCurrenciesMap(),
	}
	data.WriteResponse(w, r)
}

func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currency := strings.ToUpper(vars["currency"])
	currencyMap := model.NewRoot().GetCurrenciesMap()
	data := model.APIResponder{
		Code: 200,
	}
	if val, ok := currencyMap[currency]; ok {
		data.Data = val
		data.WriteResponse(w, r)
		return
	}
	data.Code = 204
	data.Data = "Currency code not exist!"
	data.WriteResponse(w, r)
}
