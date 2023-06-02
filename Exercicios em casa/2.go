package main

import (
	"fmt"
	"strings"
)

func main() {
	palavras := []string{"Olá", "mundo", "do", "Golang"}
	resultado := ConcatenarStrings(palavras)
	fmt.Println(resultado)
}

func ConcatenarStrings(slice []string) string {
	return strings.Join(slice, "")
}
