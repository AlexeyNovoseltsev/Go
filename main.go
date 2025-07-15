package main

import "fmt"

func main() {
	fmt.Println("___Конвертор кросс-курса валют___")
	amount,sourceCurrency, targetCurrency := getUserInput()
	result := calculatedExchange(amount, sourceCurrency, targetCurrency)
	//const EURToRUB = USDToRUB / USDToEUR

	fmt.Print("Результат конвертации: %f\n", result)
}

func getUserInput () (float64, string, string) {
	var amount float64
	var sourceCurrency string
	var targetCurrency string

	fmt.Print("Ведите сумму для конвертации: ")
	fmt.Scan(&amount)
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&sourceCurrency)
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)
	return amount, sourceCurrency, targetCurrency
}

func calculatedExchange  (amount float64, sourceCurrency, targetCurrency string) float64  {
	var result float64
	fmt.Printf("\n (Функция рассчета вызвана с параметрами: %.2f, %s, %s)\n", amount, sourceCurrency, targetCurrency)
	return result
}