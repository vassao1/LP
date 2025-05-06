package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("Insira um número inteiro: ")
	fmt.Scan(&num)

	if num < 2 {
		fmt.Println("O número deve ser maior ou igual a 2.")
		return
	}
	if num == 2 {
		fmt.Printf("%d é um número primo.\n", num)
		return
	}
	isPrimo := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrimo = false
			break
		}
	}
	if isPrimo {
		fmt.Printf("%d é um número primo.\n", num)
	} else {
		fmt.Printf("%d não é um número primo.\n", num)
	}
}
