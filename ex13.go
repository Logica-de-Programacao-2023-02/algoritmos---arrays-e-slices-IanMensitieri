package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var numeroAdicionar int
	fmt.Print("Informe um número para adicionar ao primeiro e último elemento do array: ")
	fmt.Scan(&numeroAdicionar)

	array[0] += numeroAdicionar
	array[len(array)-1] += numeroAdicionar

	fmt.Println("Array resultante:", array)
}
