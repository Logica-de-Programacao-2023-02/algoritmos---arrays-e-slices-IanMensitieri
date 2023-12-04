package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var valorProcurado int
	fmt.Print("Informe um valor para buscar no array: ")
	fmt.Scan(&valorProcurado)

	encontrado := false
	for _, elemento := range array {
		if elemento == valorProcurado {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Printf("O valor %d está presente no array.\n", valorProcurado)
	} else {
		fmt.Printf("O valor %d não está presente no array.\n", valorProcurado)
	}
}
