package main

import "fmt"

func calcularMedia(slice []int) float64 {
	total := 0
	for _, valor := range slice {
		total += valor
	}
	media := float64(total) / float64(len(slice))
	return media
}

func main() {
	numeros := []int{5, 10, 15, 20}
	media := calcularMedia(numeros)
	fmt.Printf("A média dos valores é: %.2f\n", media)
}
