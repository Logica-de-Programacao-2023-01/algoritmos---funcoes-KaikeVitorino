package main

import "fmt"

func main() {
	x := mediaAritmetica([]int{2.0, 4.0, 4.0, 6.0})
	fmt.Println(x)
}

func mediaAritmetica(slice []int) float64 {
	soma := 0.0
	for _, valor := range slice {
		soma += float64(valor)
	}
	media := soma / float64(len(slice))
	return media
}
