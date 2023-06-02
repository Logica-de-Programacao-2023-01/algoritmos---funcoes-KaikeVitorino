package main

import "fmt"

func main() {
	numeros := []int{5, 2, 9, 3, 7}
	segundoMaior := EncontrarSegundoMaior(numeros)
	fmt.Println(segundoMaior)
}

func EncontrarSegundoMaior(slice []int) int {
	if len(slice) < 2 {
		panic("Slice precisa ter pelo menos dois elementos")
	}

	maior := slice[0]
	segundoMaior := slice[1]

	for i := 2; i < len(slice); i++ {
		if slice[i] > maior {
			segundoMaior = maior
			maior = slice[i]
		} else if slice[i] > segundoMaior {
			segundoMaior = slice[i]
		}
	}

	return segundoMaior
}
