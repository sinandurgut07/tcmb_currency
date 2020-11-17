package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	currencyAPI := NewCurrencyAPI()
	e.GET("/", currencyAPI.GetAllCurrencies)
	e.GET("/currency/:code", currencyAPI.GetCurrency)
	e.Logger.Fatal(e.Start(":1323"))
}
