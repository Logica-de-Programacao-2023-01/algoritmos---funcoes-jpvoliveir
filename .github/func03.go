package main

import "fmt"

func concatenarStrings(slice []string) string {
	resultado := ""
	for _, str := range slice {
		resultado += str
	}
	return resultado
}

func main() {
	palavras := []string{"Olá", " ", "mundo!"}
	concatenacao := concatenarStrings(palavras)
	fmt.Println(concatenacao)
}
