package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64
	fmt.Print("Insira a temperatura em fahrenheit: ")
	fmt.Scan(&fahrenheit)
	celsius := ((fahrenheit - 32) * 5) / 9
	fmt.Printf("A temperatura em celsius é: %.2f\n", celsius)
}
