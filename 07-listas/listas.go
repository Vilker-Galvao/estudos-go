package main

import "fmt"

func main() {
	//Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0]) // Retorna somente o que tem dentro da string
	//fmt.Println(array[0], array[1]) // Retorna somente o que tem dentro da string
	//fmt.Println(array)

	//numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(numPrimos)
	//fmt.Println(numPrimos[0:3])
	//fmt.Println(numPrimos[2:])

	// Slices - sem tamanho definido
	//var slice []string
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "World"
	//fmt.Println(slice[0], slice[1])
	//fmt.Println(slice[0])
	//fmt.Println(slice[1])
	//fmt.Println(slice[2]) //Fica vazio porque não tem valor
	//slice[2] = "Vilker"
	//fmt.Println(slice[2]) //Retorna valor "Vilker" porque foi atribuido
	//fmt.Println(slice)
	//fmt.Println(len(slice))


	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14) // Função "append" adiciona valores
	fmt.Println(numPares)
}







// LISTAS
// 1 - Arrays e Slices: Homogêneos
// todos os elementos tem os mesmo tipo
// [1, 2, 3, 4, 5, 6] - []int
// ["vilker", "bento", "golang"] - []string

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro
// mao[string]int
// { "vilker": 20, "bento": 4}
// map [string]string
// { "vilker": "galvao", "bento": "galvao" }

// Array

// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do array
// Por conto da tamanho fixo, não é tão usado. Só em casos específicos

// Slice
//
// Tipo o array, mas sem tamanho fixo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do slice
// Função append() usada para adicionar valores.
