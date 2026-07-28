package main

import(
	"fmt"
	"strconv"
	"strings"
) 

func main(){
	fmt.Println("Digite sua operação (use o formato: 2*2): ")

	var input string

	fmt.Scan(&input)

	operation := strings.Split(input, "")

	result := getResult(operation)

	fmt.Printf("%s %s %s = %d\n", operation[0], operation[1], operation[2], result)

}

func getResult(operation []string) int{

	num1, _ := strconv.Atoi(operation[0])

	num2, _ := strconv.Atoi(operation[2])

	switch operation[1]{
	case "+":
		return num1 + num2

	case "-":
		return num1 - num2

	case "*":
		return num1 * num2

	case "/":
		return num1 / num2

	default:
		panic("Operação inválida!")
	}
}