package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Bool: %v\n", b)
	fmt.Printf("String: %v\n", s)
	i = 11
	fmt.Printf("Inteiro: %v\n", i)


	// Se não declarar valores, o GO declara "zero values"
}