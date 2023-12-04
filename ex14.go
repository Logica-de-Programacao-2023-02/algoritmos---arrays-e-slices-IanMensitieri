package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80}

	var indice1, indice2 int
	fmt.Print("Informe o primeiro índice de elemento a ser trocado: ")
	fmt.Scan(&indice1)
	fmt.Print("Informe o segundo índice de elemento a ser trocado: ")
	fmt.Scan(&indice2)

	if indice1 >= 0 && indice1 < len(slice) && indice2 >= 0 && indice2 < len(slice) {
		slice[indice1], slice[indice2] = slice[indice2], slice[indice1]
		fmt.Println("Slice resultante:", slice)
	} else {
		fmt.Println("Índices fora dos limites do slice.")
	}
}
