package main

import "fmt"

func main() {
	array := [6]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	var numeroAdicionar float64
	fmt.Print("Informe um número para adicionar em todas as posições do array: ")
	fmt.Scan(&numeroAdicionar)

	for i := 0; i < len(array); i++ {
		array[i] += numeroAdicionar
	}

	fmt.Println("Array resultante:", array)
}
