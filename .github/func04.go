package main

import (
	"fmt"
	"sort"
)

func segundoMaiorValor(slice []int) int {
	// Ordena o slice em ordem decrescente
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	// Retorna o segundo elemento do slice (segundo maior valor)
	return slice[1]
}

func main() {
	numeros := []int{5, 10, 15, 20, 25}
	segundoMaior := segundoMaiorValor(numeros)
	fmt.Println("O segundo maior valor é:", segundoMaior)
}
