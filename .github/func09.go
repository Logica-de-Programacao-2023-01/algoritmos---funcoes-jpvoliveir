package main

import (
	"errors"
	"fmt"
	"strings"
)

func extrairPalavras(str string) ([]string, error) {
	if str == "" {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Split(str, " ")
	return palavras, nil
}

func main() {
	frase := "Olá mundo! Bem-vindo ao Go."
	palavras, err := extrairPalavras(frase)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(palavras)
}
