package main

import (
	"errors"
	"fmt"
)

func contemNumero(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numero := 3
	slice := []int{1, 2, 3, 4, 5}
	encontrado, err := contemNumero(numero, slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Número encontrado:", encontrado)
}
