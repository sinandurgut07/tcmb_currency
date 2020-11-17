package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/sinandurgut07/tcmb_currency/utils"
)

type CurrencyAPI struct {
	APIKey string
	Cache *cache.Cache
}

func NewCurrencyAPI() *CurrencyAPI {
	return &CurrencyAPI{
		APIKey: utils.GetEnvVars("CURRENCY_API_KEY"),
		Cache: cache.New(30*time.Second, 5*time.Minute),
	}
}

func (ca *CurrencyAPI) GetAllCurrencies(c echo.Context) error {
	resp, err := utils.GetOrSetCurrencies(ca.Cache)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (ca *CurrencyAPI) GetCurrency(c echo.Context) error {
	currencyCode := strings.ToUpper(c.Param("code"))
	if len(currencyCode) < 3 {
		return c.NoContent(http.StatusBadRequest)
	}

	resp, err := utils.GetOrSetCurrencies(ca.Cache)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	var currencyItem utils.Currency
	var found bool
	for _, item := range resp.Currency {
		if item.CurrencyCode == currencyCode {
			found = true
			currencyItem = item
		}
		continue
	}
	if found {
		return c.JSON(http.StatusOK, currencyItem)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "invalid currency code"})
} 