package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	numeros := []int{5, 2, 9, 3, 7}
	ordenados, err := OrdenarNumeros(numeros)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ordenados)
}

func OrdenarNumeros(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazio")
	}

	ordenados := make([]int, len(slice))
	copy(ordenados, slice)
	sort.Ints(ordenados)

	return ordenados, nil
}
