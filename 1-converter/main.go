package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("__Конвертер валют__")
	showResult(getResult())
}

func showResult(result float64) {
	fmt.Printf("Результат: %.2f\n", result)
}

func getCurrency() *map[string]float64 {
	var currency = map[string]float64{
		"USD": 1,
		"EUR": 0.9,
		"RUB": 80,
	}
	return &currency
}

func getCurrencyListString(currencyNames *[]string) string {
	var result string
	for _, currencyName := range *currencyNames {
		if len(result) == 0 {
			result = currencyName
		} else {
			result = result + ", " + currencyName
		}
	}

	return result
}

func filteredCurrencyList(currencyInput ...string) []string {
	var result []string
	currencyNames := getAllCurrencyName()
	for _, currencyName := range currencyNames {
		var isTrue bool = false
		for _, input := range currencyInput {
			if input == currencyName {
				isTrue = true
				break
			}
		}
		if !isTrue {
			result = append(result, currencyName)
		}
	}
	return result
}

func getAllCurrencyName() []string {
	var result []string
	for key := range *getCurrency() {
		result = append(result, key)
	}
	return result
}

func checkCurrency(currency string, currencies *[]string) (bool, error) {
	for _, currencyName := range *currencies {
		if currencyName == currency {
			return true, nil
		}
	}
	return false, errors.New("вы ввели неверное значение")
}

func getUserData[T string | float64]() T {
	var userData T
	fmt.Scan(&userData)
	return userData
}

func getResult() float64 {
	var step uint8 = 0
	var initialCurrency, targetCurrency string
	var value float64

	currencies := filteredCurrencyList()
	currenciesString := getCurrencyListString(&currencies)

	for step <= 2 {
		if step == 0 {
			fmt.Printf("Выбор исходной валюты: %v\n", currenciesString)
			initialCurrency = getUserData[string]()

			_, err := checkCurrency(initialCurrency, &currencies)

			if err != nil {
				fmt.Println(err)
				continue
			}
			step++
		}

		if step == 1 {
			fmt.Println("Сумма валюты:")
			value = getUserData[float64]()

			_, err := isNumber(value)

			if err != nil {
				fmt.Println(err)
				continue
			}
			step++
		}

		if step == 2 {
			currencies = filteredCurrencyList(initialCurrency)
			currenciesString = getCurrencyListString(&currencies)

			fmt.Printf("Выбор целевой валюты: %v\n", currenciesString)
			targetCurrency = getUserData[string]()

			_, err := checkCurrency(targetCurrency, &currencies)

			if err != nil {
				fmt.Println(err)
				continue
			}
			step++
		}
	}

	return converterCalc(initialCurrency, targetCurrency, value)
}

func converterCalc(initialCurrency string, targetCurrency string, value float64) float64 {
	currencies := getCurrency()
	var quantity float64 = (*currencies)[targetCurrency] / (*currencies)[initialCurrency]
	return quantity * value
}

func isNumber(num float64) (bool, error) {
	if num > 0 && !math.IsNaN(num) {
		return true, nil
	}
	return false, errors.New("Ошибка: ввели не число")
}
