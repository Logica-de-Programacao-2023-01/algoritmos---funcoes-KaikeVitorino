package main

import (
	"errors"
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6}
	pares, err := FiltrarPares(numeros)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pares)
}

func FiltrarPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazio")
	}

	pares := make([]int, 0)
	for _, valor := range slice {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}

	return pares, nil
}
