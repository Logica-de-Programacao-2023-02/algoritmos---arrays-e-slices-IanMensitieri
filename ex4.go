package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Digite o tamanho do Slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Digite o valor para o elemento %d: ", i+1)
		fmt.Scanln(&slice[i])
	}

	fmt.Printf("O Slice criado é: %v\n", slice)

	var soma int
	for _, valor := range slice {
		soma += valor
	}

	fmt.Printf("A soma dos valores no Slice é: %d\n", soma)
}
