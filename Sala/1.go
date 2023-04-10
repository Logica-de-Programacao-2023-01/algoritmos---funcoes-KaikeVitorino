package main

import "errors"
import "fmt"

func main() {
	var a, b int
	fmt.Println("Digite a:")
	fmt.Scan(&a)
	fmt.Println("Digite b:")
	fmt.Scan(&b)
	fmt.Println(Divisao(a, b))
}

func Divisao(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Erro: Divisão por zero nao existe, seu animal")
	}
	resultado := float64(a) / float64(b)
	return resultado, nil

}
