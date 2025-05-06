package main

import (
	"fmt"
)

func main() {
	var impar int
	fmt.Print("Insira quantos ímpares você quer imprimir: ")
	fmt.Scan(&impar)
	if impar < 0 {
		fmt.Println("O número deve ser um inteiro não negativo.")
		return
	}
	cont := 0
	for i := 0; cont < impar; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
			cont++
		}
	}
	fmt.Println()
}
