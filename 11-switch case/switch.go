package main

import (
	"fmt"
)

func main() {

	posicao := 2
	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	}

	nome := "bento"
	switch nome {
	case "vilker":
		fmt.Println("É uma pessoa")
	case "bento":
		fmt.Println("É um cachorro")
	case "bob":
		fmt.Println("É um personagem")
	}
}