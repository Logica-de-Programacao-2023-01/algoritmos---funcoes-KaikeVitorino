package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	frase := "Olá mundo do Golang"
	palavras, err := ExtrairPalavras(frase)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(palavras)
}

func ExtrairPalavras(frase string) ([]string, error) {
	if frase == "" {
		return nil, errors.New("String está vazia")
	}

	palavras := strings.Split(frase, " ")
	return palavras, nil
}
