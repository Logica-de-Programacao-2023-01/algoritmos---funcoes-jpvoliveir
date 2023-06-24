package main

import (
	"errors"
	"fmt"
)

type Operacao func(int) int

func aplicarFuncao(slice []int, funcao Operacao) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]int, len(slice))
	for i, valor := range slice {
		resultado[i] = funcao(valor)
	}

	return resultado, nil
}

func dobrarNumero(numero int) int {
	return numero * 2
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	resultado, err := aplicarFuncao(numeros, dobrarNumero)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
