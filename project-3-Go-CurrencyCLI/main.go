package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"currency-converter/currency"
  "currency-converter/config"
)

func main() {
cnfig := config.LoadConfig("config.json")

	data, err := currency.FetchData(cnfig.APIKey)
	if err != nil {
		fmt.Println("Error fetching currency exchange rates data:", err)
		os.Exit(1)
	}	

  reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the amount to convert: ")
	amountStr, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(amountStr[:len(amountStr)-1], 64)
	if err != nil {
		fmt.Println("Invalid amount entered:", err)
		os.Exit(1)
	}

	fmt.Print("Enter the currency code to convert from (e.g. USD, EUR, etc.): ")
	fromCode, _ := reader.ReadString('\n')
	fromCode = fromCode[:len(fromCode)-1]
	if _, ok := currency.Currencies[fromCode]; !ok {
		fmt.Println("Invalid currency code entered:", fromCode)
		os.Exit(1)
	}

	fmt.Print("Enter the currency code to convert to (e.g. USD, EUR, etc.): ")
	toCode, _ := reader.ReadString('\n')
	toCode = toCode[:len(toCode)-1]
	if _, ok := currency.Currencies[toCode]; !ok {
		fmt.Println("Invalid currency code entered:", toCode)
		os.Exit(1)
	}

	converted := data.ConvertAmount(amount, fromCode, toCode)
	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, currency.Currencies[fromCode], converted, currency.Currencies[toCode])
}

