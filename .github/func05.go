package main

import "fmt"

func encontrarPosicao(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	numeros := []int{5, 10, 15, 20, 25}
	valor := 15
	posicao := encontrarPosicao(numeros, valor)
	fmt.Printf("O valor %d está na posição %d\n", valor, posicao)
}
