package main

import (
	"fmt"
	"strings"
)

func main() {
	palavras := []string{"Olá", "mundo", "do", "Golang"}
	resultado := concatenarStrings(palavras)
	fmt.Println(resultado)
}

func concatenarStrings(slice []string) string {
	return strings.Join(slice, "")
}
