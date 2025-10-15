package main

import "fmt"

func main() {
	const USD = 1
	const EUR = 0.9
	const RUB = 80
	const eurToRub = RUB / EUR

	var quantity float64 = getUserData[float64]()
	var currencyInput string = getUserData[string]()
	var currencyOutput string = getUserData[string]()

	var result float64 = getResult(quantity, currencyInput, currencyOutput)
	fmt.Println(result)
}

func getUserData[T string | float64]() T {
	var userData T
	fmt.Scan(&userData)
	return userData
}

func getResult(quantity float64, currencyInput string, currencyOutput string) float64 {
	return 0.0
}
