package main

import "fmt"

func main() {
	fmt.Println(soma(42, 13))
	
	SomaDosValores := soma (10, 13)      // funciona do mesmo jeito
	fmt.Println(SomaDosValores)

	sub := subtracao(10, 5)
	fmt.Println(sub)

	nome, sobrenome := printNomeCompleto("VILKER", "GALVAO")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}

// Função começando com letra minúscula:
// FUnção é PRIVADA
// Função ela só pode ser utilizada no próprio pacote
func printNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// Função começando com letra maiúscula:
// FUnção é PÚBLICA
// Função pode ser utilizada fora do próprio pacote
// Como utilizaria ela fora: main.PrintNome(nome)
func PrintNome(nome string) string {
	return nome
}



func subtracao (x int, y int) int {
	return x - y
}

func soma (x int, y int) int {
	return x + y
}