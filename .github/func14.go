package main

import (
	"errors"
	"fmt"
)

func intersecao(slice1 []int, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	mapa := make(map[int]bool)
	intersecao := make([]int, 0)

	for _, num := range slice1 {
		mapa[num] = true
	}

	for _, num := range slice2 {
		if mapa[num] {
			intersecao = append(intersecao, num)
		}
	}

	return intersecao, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}
	resultado, err := intersecao(slice1, slice2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
