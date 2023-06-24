package main

import (
	"errors"
	"fmt"
)

func numerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	pares := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}

	return pares, nil
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6}
	pares, err := numerosPares(numeros)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pares)
}
