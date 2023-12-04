package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "a", "d", "e", "a", "f"}

	var valorRemover string
	fmt.Print("Informe um valor para remover do slice: ")
	fmt.Scan(&valorRemover)

	novoSlice := removerOcorrencias(slice, valorRemover)

	fmt.Println("Slice resultante:", novoSlice)
}

func removerOcorrencias(slice []string, valor string) []string {
	resultado := make([]string, 0, len(slice))

	for _, elemento := range slice {
		if elemento != valor {
			resultado = append(resultado, elemento)
		}
	}

	return resultado
}
