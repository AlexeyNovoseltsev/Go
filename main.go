package main

import "fmt"

func main() {
	const USDToEUR =  0.86
	const USDToRUB = 78.00
	const EURToRUB =  USDToRUB / USDToEUR
	fmt.Print(EURToRUB)
}