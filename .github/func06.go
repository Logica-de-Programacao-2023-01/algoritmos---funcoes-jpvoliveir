package main

import (
	"errors"
	"fmt"
	"strings"
)

func concatenarComVirgula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	resultado := strings.Join(slice, ",")
	return resultado, nil
}

func main() {
	palavras := []string{"Olá", "mundo", "Go"}
	concatenacao, err := concatenarComVirgula(palavras)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(concatenacao)
}
