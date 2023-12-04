package main

import "fmt"

func main() {
	array := lerArray("Informe os elementos do array:", -1)

	if estaOrdenadoCrescente(array) {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}

func lerArray(mensagem string, tamanho int) []int {
	fmt.Println(mensagem)
	var n int
	fmt.Print("Informe o tamanho do array: ")
	fmt.Scan(&n)

	if tamanho != -1 && n != tamanho {
		fmt.Printf("Por favor, informe um array com tamanho %d.\n", tamanho)
		return lerArray(mensagem, tamanho)
	}

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&array[i])
	}
	return array
}

func estaOrdenadoCrescente(array []int) bool {
	n := len(array)

	for i := 1; i < n; i++ {
		if array[i] < array[i-1] {
			return false
		}
	}

	return true
}
