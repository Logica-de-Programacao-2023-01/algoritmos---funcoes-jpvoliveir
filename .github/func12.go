package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func filtrarStringsMaiusculas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	var resultado strings.Builder

	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			resultado.WriteString(str)
			resultado.WriteString(" ")
		}
	}

	return strings.TrimSpace(resultado.String()), nil
}

func main() {
	palavras := []string{"Gato", "cachorro", "Rato", "Elefante"}
	resultado, err := filtrarStringsMaiusculas(palavras)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
