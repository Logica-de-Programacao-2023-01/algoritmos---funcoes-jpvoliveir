package main

import (
	"fmt"
	"strings"
)

func contarVogais(str string) int {
	vogais := "aeiouAEIOU"
	qtdVogais := 0

	for _, char := range str {
		if strings.ContainsRune(vogais, char) {
			qtdVogais++
		}
	}

	return qtdVogais
}

func main() {
	frase := "Olá, mundo!"
	qtd := contarVogais(frase)
	fmt.Printf("A quantidade de vogais na frase é: %d\n", qtd)
}
