package main

import "fmt"

func main() {
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var indiceLinha, indiceColuna int
	fmt.Print("Informe o índice de linha: ")
	fmt.Scan(&indiceLinha)
	fmt.Print("Informe o índice de coluna: ")
	fmt.Scan(&indiceColuna)

	if indiceLinha >= 0 && indiceLinha < 2 && indiceColuna >= 0 && indiceColuna < 3 {
		fmt.Printf("Valor na posição (%d, %d): %d\n", indiceLinha, indiceColuna, matriz[indiceLinha][indiceColuna])
	} else {
		fmt.Println("Índices fora dos limites da matriz.")
	}
}
