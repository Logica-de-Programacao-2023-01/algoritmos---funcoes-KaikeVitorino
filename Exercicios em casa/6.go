package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	palavras := []string{"Olá", "mundo", "do", "Golang"}
	resultado, err := ConcatenarComVirgulas(palavras)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}

func ConcatenarComVirgulas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice está vazio")
	}

	return strings.Join(slice, ","), nil
}
