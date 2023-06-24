package main

import (
	"errors"
	"fmt"
)

func stringsComMaisDe5Caracteres(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]string, 0)

	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}

func main() {
	slice := []string{"gato", "cachorro", "elefante", "cobra", "pássaro"}
	resultado, err := stringsComMaisDe5Caracteres(slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
