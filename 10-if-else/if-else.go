package main

import "fmt"

// IF / ELSE
// SE / SENÃO

func main() {

	//numero := 3
	// if (condição) {ação}
	//if numero == 1 { // true
	//	fmt.Println("Valor é igual a 1")
	//} else {
	//	fmt.Println("Valor não é igual a 1")
	//}

	// if numero == 1 { // true
	// 	fmt.Println("Valor é igual a 1")
	// } else if numero == 2 {
	// 	fmt.Println("Valor é igual a 2")
	// } else {
	// 	fmt.Println("Valor é diferente de 1 e 2")
	// }

	numero := 8
	if numero%2 == 0 {
		fmt.Printf("%d é par\n", numero)
	} else {
		fmt.Printf("%d é impar\n", numero)
	}

	nome := "vilker"
	if nome == "vilker" {
		fmt.Println("Acetou! =)")
	
	} else {
		fmt.Println("Errou! =(")
	}

	// fmt.Println(1 + 1)
	// fmt.Println(2 - 1)
	// fmt.Println(2 * 2)
	// fmt.Println(10 / 2)
	// fmt.Println(10 % 2)
}

// Operações
// Soma: 1 + 1 
// Subtração: 2 - 1
// Divisão: 10 / 2 
// Multiplicação: 2 * 2
// Resto da divisão por x: 7%2 (resto da divisão por 2)