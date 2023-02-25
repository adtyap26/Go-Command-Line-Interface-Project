package currency

import (
	"encoding/json"
	"net/http"

  "currency-converter/config"
)

const baseURL = "https://openexchangerates.org/api/latest.json?app_id="

var Currencies = map[string]string{
	"AUD": "Australian dollar",
	"BRL": "Brazilian real",
	"CAD": "Canadian dollar",
	"CHF": "Swiss franc",
	"CLP": "Chilean peso",
	"CNY": "Chinese yuan",
	"CZK": "Czech koruna",
	"DKK": "Danish krone",
	"EUR": "Euro",
	"GBP": "Pound sterling",
	"HKD": "Hong Kong dollar",
	"HUF": "Hungarian forint",
	"IDR": "Indonesian rupiah",
	"ILS": "Israeli shekel",
	"INR": "Indian rupee",
	"JPY": "Japanese yen",
	"KRW": "South Korean won",
	"MXN": "Mexican peso",
	"MYR": "Malaysian ringgit",
	"NOK": "Norwegian krone",
	"NZD": "New Zealand dollar",
	"PHP": "Philippine peso",
	"PKR": "Pak Pakistani rupee",
	"PLN": "Polish zloty",
	"RUB": "Russian ruble",
	"SEK": "Swedish krona",
	"SGD": "Singapore dollar",
	"THB": "Thai baht",
	"TRY": "Turkish lira",
	"TWD": "New Taiwan dollar",
	"USD": "United State dollar",
	"ZAR": "South African rand",
}

// Currency holds the currency information
type Currency struct {
	Rates map[string]float64 `json:"rates"`
}

// FetchData fetches the latest currency exchange rates data from the API
func FetchData(apikkey string) (*Currency, error) {
  // appID := config.APIKey
  cnfig := config.LoadConfig("config.json")
  url := baseURL + cnfig.APIKey

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data Currency
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

// ConvertAmount converts an amount from one currency to another
func (c *Currency) ConvertAmount(amount float64, fromCode, toCode string) float64 {
	fromRate := c.Rates[fromCode]
	toRate := c.Rates[toCode]
	converted := amount / fromRate * toRate

	return converted
}

