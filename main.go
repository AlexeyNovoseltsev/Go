package main

import (
	"fmt"
	"strings"
)

const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
)

func main() {

	showGreeting()
	// Бесконечный цикл для повторного использования конвертора. Пока пользователь не откажется от продолжения цикл будет выполняться.
	for {
		amount, sourceCurrency, targetCurrency := getUserInput()
		result := calculateExchange(amount, sourceCurrency, targetCurrency)
		printResult(amount, sourceCurrency, result, targetCurrency)
		if !askToContinue() {
			break
		}
	}

}

func askToContinue() bool {
	for {
		fmt.Print("\nХотите выполнить еще один расчет? (Y/N): ")
		var answer string
		fmt.Scan(&answer)
		if answer == "Y" {
			return true
		} else if answer == "N" {
			return false
		}
		fmt.Println("Некорректный ввод. Пожалуйста, введите Y или N.")
	}
}

func getUserInput() (float64, string, string) {
	// Шаг 1: Получаем исходную валюту, вызывая нашу универсальную функцию.
	sourceCurrency := getCurrency("Введите исходную валюту (USD, EUR, RUB): ", "")
	// Шаг 2: Получаем сумму для конвертации.
	amount := getAmount()
	// Шаг 3: Передаем исходную валюту
	targetCurrency := getCurrency("Введите целевую валюту (USD, EUR, RUB): ", sourceCurrency)
	// Возвращаем проверенные значения.
	return amount, sourceCurrency, targetCurrency

}

// Получение и проверка валюты
func getCurrency(prompt, forbiddenCurrency string) string {
	var currency string
	// Бесконечный цикл, который прервется только после корректного ввода.
	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)
		// Проверка: Валюта входит в список разрешенных?
		if currency != USD && currency != EUR && currency != RUB {
			fmt.Println("Ошибка: Неверная валюта. Пожалуйста, выберите USD, EUR, RUB.")
			// Проверка: Валюта не совпадает с запрещенной?
		} else if currency == forbiddenCurrency {
			fmt.Println("Ошибка: Исходная и целевая валюты не могут совпадать.")
		} else {
			return currency
		}
	}
}

// Получение и проверка суммы
func getAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму для конвертации: ")
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Ошибка: Введите положительное число.")
			var discard string
			fmt.Scanln(&discard)
		} else {
			return amount
		}
	}
}

func showGreeting() {
	fmt.Println("___Конвертер кросс-курса валют___")
}

func printResult(amount float64, sourceCurrency string, result float64, targetCurrency string) {
	// %.2f - специальный формат для вывода числа с 2 знаками после запятой.
	fmt.Printf("\nРезультат конвертации: %.2f %s = %.2f %s\n", amount, sourceCurrency, result, targetCurrency)
}

func calculateExchange(amount float64, sourceCurrency, targetCurrency string) float64 {
	fmt.Printf("\n(Функция расчета вызвана с параметрами: %.2f, %s, %s)\n", amount, sourceCurrency, targetCurrency)
	var result float64

	switch {
	case sourceCurrency == USD && targetCurrency == EUR:
		result = amount * 0.85
	case sourceCurrency == USD && targetCurrency == RUB:
		result = amount * 79.55
	case sourceCurrency == EUR && targetCurrency == USD:
		result = amount * 1.17
	case sourceCurrency == EUR && targetCurrency == RUB:
		result = amount * 93.35
	case sourceCurrency == RUB && targetCurrency == USD:
		result = amount / 79.55
	case sourceCurrency == RUB && targetCurrency == EUR:
		result = amount / 93.35
	default:
		fmt.Println("Ошибка: Неизвестная комбинация валют!")
		result = 0
	}
	return result
}
