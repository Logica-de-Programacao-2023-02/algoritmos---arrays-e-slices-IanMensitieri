package main

import "fmt"

func main() {
	array := [3]int{2, 4, 6}

	var soma int
	for _, valor := range array {
		soma += valor
	}

	fmt.Printf("A soma dos valores no array é: %d\n", soma)
}
