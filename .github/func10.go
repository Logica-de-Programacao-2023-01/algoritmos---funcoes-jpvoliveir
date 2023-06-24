package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarCrescente(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	sliceOrdenado := make([]int, len(slice))
	copy(sliceOrdenado, slice)

	sort.Ints(sliceOrdenado)
	return sliceOrdenado, nil
}

func main() {
	numeros := []int{4, 2, 7, 1, 5}
	ordenados, err := ordenarCrescente(numeros)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ordenados)
}
