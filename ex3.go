package main

import "fmt"

func main() {
	array := [4]float64{2.0, 3.0, 1.5, 4.0}

	produto := 1.0
	for _, valor := range array {
		produto *= valor
	}

	fmt.Printf("O produto dos valores no array é: %.2f\n", produto)
}
