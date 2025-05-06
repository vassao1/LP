package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insira a string principal: ")
	mainstr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a string principal:", err)
		return
	}
	mainstr = strings.TrimSpace(mainstr)
	fmt.Print("Insira a substring: ")
	substr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a substring:", err)
		return
	}
	substr = strings.TrimSpace(substr)

	if len(substr) == 0 || len(mainstr) == 0 {
		fmt.Println("As strings não podem ser vazias.")
		return
	}
	if len(substr) > len(mainstr) {
		fmt.Println("A substring não pode ser maior que a string principal.")
		return
	}
	cont := 0
	for i := 0; i <= len(mainstr)-len(substr); i++ {
		if mainstr[i] == substr[0] {
			j := 0
			for j = 0; j < len(substr); j++ {
				if mainstr[i+j] != substr[j] {
					break
				}
			}
			if j == len(substr) {
				cont++
			}
		}
	}
	fmt.Printf("A substring '%s' aparece %d vez(es) na string principal '%s'.\n", substr, cont, mainstr)
}
