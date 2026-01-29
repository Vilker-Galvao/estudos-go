package main

import "fmt"

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes

type Pessoa struct {
	Nome string
	Idade int
}
func main() {
	fmt.Println(Pessoa{"Vilker", 20})
	fmt.Println(Pessoa{Nome: "Bento", Idade: 4})
	fmt.Println(Pessoa{Nome: "Vilker"})

	p1 := Pessoa{Nome: "Bob", Idade: 2}
	//fmt.Println(p1.Nome)
	//fmt.Println(p1.Idade)

	p1.Idade = 3
	//fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade: 5}
	
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	//fmt.Println(pessoas)

	//structs + map
	//alunos := map[string][]Pessoa{}
	//alunos["programação"] = pessoas
	//fmt.Println(alunos)

	var alunos = map[string][]Pessoa{
		"Programação": {{Nome: "Maria", Idade: 19}, {Nome: "Bento", Idade: 21}},
		"Engenharia": {{Nome: "Marcos", Idade: 19}, {Nome: "Ana", Idade: 21}},
	}
	fmt.Println(alunos)
}
