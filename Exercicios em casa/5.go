package main

import "fmt"

func main() {
	numeros := []int{3, 7, 5, 2, 9}
	posicao := EncontrarPosicaoElemento(numeros, 5)
	fmt.Println(posicao)
}

func EncontrarPosicaoElemento(slice []int, valor int) int {
	for i, v := range slice {
		if v == valor {
			return i
		}
	}

	return -1
}
