package main

import (
	"fmt"
)

func main() {
	var base, exponente int
	fmt.Print("Insira a base: ")
	fmt.Scan(&base)
	fmt.Print("Insira o expoente: ")
	fmt.Scan(&exponente)
	if exponente < 0 {
		fmt.Println("O expoente deve ser um número inteiro não negativo.")
		return
	}
	resultado := base
	for i := 1; i < exponente; i++ {
		resultado *= base
	}
	fmt.Printf("%d elevado a %d é igual a %d\n", base, exponente, resultado)
}
