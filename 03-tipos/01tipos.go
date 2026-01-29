package main

import (
	"fmt"
)

func main() {
	// bool (true / false)
	// fmt.Printlf("type: %T - Value: %v\n", false, false)

	// string - sequência de bytes
	fmt.Printf("Type: %T - Value: %v\n", "Vilker", "Vilker")
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")
	
	// int
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	
	//float - decimal
	fmt.Printf("Type: %T - Value: %v\n", 1.123, 1.123)
}

// Tipos:
// bool (true / false)
// string - sequência de bytes
// int - números inteiros
// float (float64, float32) - para números decimais