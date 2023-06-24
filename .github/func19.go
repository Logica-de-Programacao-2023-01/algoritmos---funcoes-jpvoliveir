package main

import (
	"errors"
	"fmt"
)

func isPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func primosAte(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("O número é negativo")
	}

	primos := make([]int, 0)

	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func main() {
	numero := 20
	resultado, err := primosAte(numero)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
