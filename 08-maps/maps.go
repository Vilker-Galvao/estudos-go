package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["vilker"] = 20
	idade["bento"] = 12
	fmt.Println(idade)
	fmt.Println(idade["vilker"])
	fmt.Println(idade["bento"])

	anoNasc := map[string]int{
		"vilker": 2005,
		"bento": 2014,
	}
	fmt.Println(anoNasc)

	
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["vilker"])
	fmt.Println(anoNasc["bento"])
	fmt.Println(anoNasc["golangDoZero"])
}





// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro tipo ou o mesmo
// map[k]v -> k = chave, v = valor

// map[string]int
// {"vilker": 20, "bento": 4}
// map[string]string
// { "vilker": "galvao", "bento": "galvao"}.
